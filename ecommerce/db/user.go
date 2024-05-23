package db

// User struct represents the user table
type User struct {
	Username string `db:"username"    validate:"required,min=4,max=12" json:"username"`
	Password string `db:"password"    validate:"required,min=4,max=8"  json:"password"`
	Email    string `db:"email"       validate:"required,email"        json:"email"   `
}

// Login creadintials are taken here
type Login struct {
	Username string `db:"username"    validate:"required,min=4,max=12" json:"username"`
	Password string `db:"password"    validate:"required,min=4,max=8"  json:"password"`
}

func InsertNewUser(newUser User) error {

	db := GetDB()
	_, err := db.Exec(`INSERT INTO "user" (username, password, email) VALUES ($1, $2, $3)`, newUser.Username, newUser.Password, newUser.Email)
	return err
}

func LoginUser(userLogin Login) error {

	db := GetDB()
	var username string
	err := db.Get(&username, `SELECT username FROM "user" WHERE username = $1 AND password = $2`, userLogin.Username, userLogin.Password)

	return err
}
