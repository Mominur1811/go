package handlers

import (
	"gobasic/web/message"
	"net/http"
)

func View(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		message.SendData(w, users)
		return
	}
	http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
}
