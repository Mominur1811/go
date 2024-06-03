package db

import (
	"ecommerce/logger"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

// User struct represents the user table
type User struct {
	Username string `db:"username"    validate:"required,min=4,max=12" json:"username"`
	Password string `db:"password"    validate:"required,min=4,max=8"  json:"password"`
	Email    string `db:"email"       validate:"required,email"        json:"email"   `
}

type UserRegisterRepo struct {
	UserTableName string
}

var userRegisterRepo *UserRegisterRepo

func InitRegisterRepo() {
	userRegisterRepo = &UserRegisterRepo{UserTableName: `"user"`}
}

func GetUserRegisterRepo() *UserRegisterRepo {
	return userRegisterRepo
}

// Login creadintials are taken here
type Login struct {
	Username string `db:"username"    validate:"required,min=4,max=12" json:"username"`
	Password string `db:"password"    validate:"required,min=4,max=8"  json:"password"`
}

type UserRepo struct {
	UserTableName string
}

var userRepo *UserRepo

func GetUserLoginRepo() *UserRepo {

	return userRepo
}

func (r *UserRegisterRepo) InsertNewUser(newUser *User) (*User, error) {

	column := map[string]interface{}{
		"username": newUser.Username,
		"password": newUser.Password,
		"email":    newUser.Email,
	}
	var columns []string
	var values []any
	for columnName, columnValue := range column {

		columns = append(columns, columnName)
		values = append(values, columnValue)

	}
	qry, args, err := GetQueryBuilder().
		Insert(r.UserTableName).
		Columns(columns...).
		Suffix(`
			RETURNING 
			username,
			password,
			email
		`).
		Values(values...).
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create audit log insert query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
		return nil, err
	}
	// Execute the SQL query and get the result
	var insertedUser User
	err = GetReadDB().QueryRow(qry, args...).Scan(&insertedUser.Username, &insertedUser.Password, &insertedUser.Email)
	if err != nil {
		slog.Error(
			"Failed to execute insert query",
			logger.Extra(map[string]interface{}{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
		return nil, err
	}

	return &insertedUser, nil

}

func (r *UserRepo) FindUser(userLogin *Login) error {
	db := GetReadDB() // Assuming GetReadDB() returns a *sql.DB

	// Build the query using squirrel
	query := GetQueryBuilder().Select("username").From("user").
		Where(sq.Eq{"username": userLogin.Username, "password": userLogin.Password})

	// Get the generated SQL and arguments from the squirrel query builder
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	// Execute the query
	var username string
	err = db.Get(&username, sql, args...)
	return err
}
