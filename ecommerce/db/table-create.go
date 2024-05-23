package db

import "fmt"

func CreateTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS "user" (
		username TEXT PRIMARY KEY,
		password TEXT,
		email TEXT
	);
	`

	db := GetDB()
	if _, err := db.Exec(createUserTable); err != nil {
		panic(err)
	}
	fmt.Println("User Table created successfully")
	createProductTable := `
	CREATE TABLE IF NOT EXISTS product (
		product_id          SERIAL PRIMARY KEY,
		product_name        TEXT,
		product_category    TEXT,
		product_price       NUMERIC,
		product_quantity    INTEGER
	);
	`

	if _, err := db.Exec(createProductTable); err != nil {
		panic(err)
	}

	fmt.Println("Product Table created successfully")

	createOrderTable := `
	CREATE TABLE IF NOT EXISTS "order" (
		username    TEXT,
		product_id       INTEGER,
		order_date       DATE,
		order_quantity   INTEGER,
		total_pay        NUMERIC,
		PRIMARY KEY (username, product_id),
		FOREIGN KEY (username) REFERENCES "user"(username),
		FOREIGN KEY (product_id) REFERENCES product(product_id)
	);
	`

	if _, err := db.Exec(createOrderTable); err != nil {
		panic(err)
	}

	fmt.Println("Order Table created successfully")

}
