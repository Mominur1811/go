package web

import (
	"gobasic/web/handlers"
	"net/http"
)

func Initialize_routes() {

	http.HandleFunc("/create", handlers.Create)
	http.HandleFunc("/update", handlers.Update)
	http.HandleFunc("/view", handlers.View)
	http.HandleFunc("/delete", handlers.Delete)
}
