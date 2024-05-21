package db

import "fmt"

func CreateTable() {

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS user_ecommerce (
		user_id SERIAL PRIMARY KEY,
		username      TEXT,
		password      TEXT,
		contactNo   TEXT
	)`

	db := GetDB()
	if _, err := db.Exec(createTableSQL); err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")
}
