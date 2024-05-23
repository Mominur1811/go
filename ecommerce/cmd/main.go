package main

import (
	"ecommerce/db"
	"ecommerce/web"
	"log"
)

func main() {
	if err := db.Init_Db(`/home/mominur/config/config.txt`); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	db.CreateTable()

	defer db.CloseDB()

	web.StartServer()
}
