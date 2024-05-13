package web

import (
	"fmt"
	"gobasic/web/middlewires"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()

	manager := middlewires.NewManager()
	InitRoutes(mux, manager)

	fmt.Println("Server Started")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
