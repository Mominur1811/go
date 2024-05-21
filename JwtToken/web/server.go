package web

import (
	"JwtToken/web/middlewire"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()

	manager := middlewire.NewManager()
	InitRoutes(mux, manager)

	fmt.Println("Server Started")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
