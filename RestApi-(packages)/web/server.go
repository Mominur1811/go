package web

import "log"

func Start_Server() {
	Initialize_routes()
	log.Println("Server started on :8080")
}
