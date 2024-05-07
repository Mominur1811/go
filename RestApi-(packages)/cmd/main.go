package main

import (
	"gobasic/web"
	"log"
	"net/http"
)

func main() {

	web.Start_Server()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
