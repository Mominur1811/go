package handlers

import (
	"gobasic/db"
	"gobasic/web/message"
	"net/http"

	_ "github.com/lib/pq"
)

func View(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		employees, err := db.ViewTable()

		if err != nil {
			message.Send_Json(w, http.StatusExpectationFailed, err)
			return
		} else {
			message.SendData(w, employees)
			return
		}

	} else {
		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
		return
	}

}
