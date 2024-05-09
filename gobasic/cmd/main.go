package main

import (
	"gobasic/db"
	"gobasic/web"
	"log"
	"net/http"
)

func main() {

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.GetDB().Close()

	web.Start_Server()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
