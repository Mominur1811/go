package db

type User struct {
	Id       int    `db:"id" 		json:"id"`
	Name     string `db:"name" 		json:"name"`
	Password string `db:"password" 	json:"password"`
}

type UserId struct {
	ID int `db:"id" json:"id"`
}

// Insertion
func InsertUser(new_user User) error {

	db := GetDB()
	_, err := db.NamedExec("INSERT INTO employee(id, name, password) VALUES(:id, :name, :password)", new_user)
	return err

}

// Deletion by taking key
func DeleteUser(delete_user UserId) error {

	db := GetDB()
	_, err := db.NamedExec("DELETE FROM employee WHERE id=:id", delete_user)
	return err
}

// Update whole row
func UpdateUser(updateUser User) error {

	db := GetDB()
	_, err := db.NamedExec("UPDATE employee SET name=:name, password=:password WHERE id=:id", updateUser)
	return err

}

// Check if user exist or not
func CheckExistenseOfUser(find_user int) error {

	db := GetDB()
	var name string
	err := db.Get(&name, "SELECT name FROM employee WHERE id=$1", find_user)
	return err
}

// view atmost 100 rows from the table
func ViewTable() (interface{}, error) {
	db := GetDB()
	var employees []User
	err := db.Select(&employees, "SELECT id, name, crypt(password, gen_salt('md5')) AS password FROM employee LIMIT 100")
	return employees, err
}
