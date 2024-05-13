package db

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB // Database connection variable

// InitializeDB initializes the database connection
func InitializeDB() error {
	var err error
	DB, err = sqlx.Connect("postgres", "host=localhost port=5432 user=root password=admin dbname=root sslmode=disable")
	return err
}

// Closing DB Connection
func CloseDB() {
	DB.Close()
}

// GetDB returns the database connection
func GetDB() *sqlx.DB {
	return DB
}
