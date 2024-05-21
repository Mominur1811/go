package db

func InsertNewUser(newUser User) error {

	db := GetDB()
	_, err := db.NamedExec("INSERT INTO user_ecommerce(username, password, contactno) VALUES(:username, :password, :contactno)", newUser)
	return err
}

func UserLoginValidation(userLogin LoginData) error {

	db := GetDB()
	var username string
	err := db.Get(&username, "SELECT username FROM user_ecommerce WHERE username = $1 AND password = $2", userLogin.UserName, userLogin.Password)

	return err
}
