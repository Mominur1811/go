package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init_Db(configData string) error {

	configParam, err := config.ReadDBConfigFromFile(configData)
	if err != nil {
		return err
	}

	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configParam.Host, configParam.Port, configParam.User, configParam.Password, configParam.DBName, configParam.SSLMode)

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
