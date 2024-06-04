package db

import (
	"ecommerce/logger"
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

// User struct represents the user table
type User struct {
	Email    string `db:"email"       validate:"required,email"        json:"email"   `
	Username string `db:"username"    validate:"required,min=4,max=12" json:"username"`
	Password string `db:"password"    validate:"required,min=4,max=8"  json:"password"`
}

type UserRegisterRepo struct {
	UserTableName string
}

var userRegisterRepo *UserRegisterRepo

func InitRegisterRepo() {
	userRegisterRepo = &UserRegisterRepo{UserTableName: `customer`}
}

func GetUserRegisterRepo() *UserRegisterRepo {
	return userRegisterRepo
}

// Login creadintials are taken here
type Login struct {
	Email    string `db:"email"    validate:"required,min=10,max=50" json:"email"`
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
			"Failed to create insert user query",
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

func (r *UserRepo) FindUser(userLogin *Login) (string, error) {
	db := GetReadDB()

	// Build the query using squirrel
	query := GetQueryBuilder().Select("email").From(`customer`).
		Where(sq.Eq{"email": userLogin.Email, "password": userLogin.Password})

	// Get the generated SQL and arguments from the squirrel query builder
	sql, args, err := query.ToSql()
	fmt.Println(sql)
	if err != nil {
		return "", err
	}

	// Execute the query
	var email string
	err = db.Get(&email, sql, args...)
	return email, err
}

func (r *UserRepo) InsertOtp(email string, otp string) error {

	_, err := GetReadDB().Exec("INSERT INTO otp_check (email, otp) VALUES ('" + email + "', '" + otp + "')")
	return err
}

func (r *UserRepo) UpdateCheckMarkUser(email string) error {

	_, err := GetReadDB().Exec("UPDATE customer SET is_valid = TRUE WHERE email = $1", email)
	return err
}
