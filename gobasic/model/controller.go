package model

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB // Database connection variable

// InitializeDB initializes the database connection
func InitializeDB() error {
	var err error
	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=root password=admin dbname=root sslmode=disable")
	return err
}

// GetDB returns the database connection
func GetDB() *sqlx.DB {
	return db
}
