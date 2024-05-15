package db

import "fmt"

func CreateProductTable() {

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		product_id SERIAL PRIMARY KEY,
		pname      TEXT,
		price      INTEGER,
		orderdate  TIMESTAMP,
		customer   TEXT
	)
	`

	db := GetDB()
	if _, err := db.Exec(createTableSQL); err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")
}
