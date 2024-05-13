package handlers

import (
	"gobasic/db"
	"gobasic/web/message"
	"net/http"

	_ "github.com/lib/pq"
)

func View(w http.ResponseWriter, r *http.Request) {

	employees, err := db.ViewTable()

	if err != nil {
		message.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}
	message.SendData(w, employees)

}
