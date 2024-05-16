package web

import (
	"fmt"
	"log"
	"net/http"

	"example.com/go_pagination_filtering_sorting/web/middlewire"
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
