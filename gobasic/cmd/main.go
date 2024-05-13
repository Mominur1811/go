package main

import (
	"gobasic/db"
	"gobasic/web"
	"log"
)

func main() {

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	defer db.CloseDB()

	web.StartServer()

}
