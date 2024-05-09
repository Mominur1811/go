package db

type User struct {
	ID       int    `db:"id" 		json:"id"`
	Name     string `db:"name" 		json:"name"`
	Password string `db:"password" 	json:"password"`
}

type UserData struct {
	ID int `db:"id" json:"id"`
}

func InsertUser(new_user User) error {

	db := GetDB()
	_, err := db.NamedExec("INSERT INTO employee(id, name, password) VALUES(:id, :name, :password)", new_user)
	return err

}

func DeleteUser(delete_user UserData) (bool, string) {
	db := GetDB()
	check, err := CheckExistenseOfUser(delete_user)
	if check {
		_, err := db.NamedExec("DELETE FROM employee WHERE id=:id", delete_user)
		if err != nil {
			return false, "Error deleting Data"
		}
		return true, "Deleted Users data"
	} else {
		if err != nil {
			return false, "Error in connecting to Database"
		}
		return false, "table has no records regarding the User ID"
	}

}

func UpdateUser(updateUser User) (bool, string) {

	db := GetDB()
	var upUser UserData
	upUser.ID = updateUser.ID
	check, err := CheckExistenseOfUser(upUser)
	if check {
		_, err := db.NamedExec("UPDATE employee SET name=:name, password=:password WHERE id=:id", updateUser)
		if err != nil {
			return false, "Error Updating Data"
		}
		return true, "Updated User data"
	} else {
		if err != nil {
			return false, "Error in connecting to Database"
		}
		return false, "Table has no records regarding the User ID"
	}
}

func ViewTable() (interface{}, error) {
	db := GetDB()
	var employees []User
	err := db.Select(&employees, "SELECT id, name, password FROM employee")
	return employees, err
}

func CheckExistenseOfUser(find_user UserData) (bool, error) {
	db := GetDB()
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM employee WHERE id=$1)", find_user.ID)
	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}
	return true, nil
}
