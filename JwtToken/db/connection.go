package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init_Db() error {

	var err error
	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)
	DB, err = sqlx.Connect("postgres", ConnStr)
	return err

}

func CloseDB() {
	DB.Close()
}

// GetDB returns the database connection
func GetDB() *sqlx.DB {
	return DB
}
