package main

import (
	"gobasic/model"
	"gobasic/web"
	"log"
	"net/http"
)

func main() {

	if err := model.InitializeDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer model.GetDB().Close()

	web.Start_Server()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
