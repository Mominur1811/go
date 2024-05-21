package main

import (
	"JwtToken/db"
	"JwtToken/web"
	"log"
)

func main() {

	if err := db.Init_Db(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	db.CreateTable()

	defer db.CloseDB()

	web.StartServer()
}
