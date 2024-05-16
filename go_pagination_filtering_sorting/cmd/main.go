package main

import (
	"log"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web"
)

func main() {
	if err := db.Init_Db(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	db.CreateProductTable()

	defer db.CloseDB()

	web.StartServer()
}
