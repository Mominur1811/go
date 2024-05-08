package handlers

import (
	"fmt"
	"gobasic/model"
	"gobasic/web/message"
	"net/http"

	_ "github.com/lib/pq"
)

func View(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		db := model.GetDB()

		// Retrieve data from the database
		var employees []User
		err := db.Select(&employees, "SELECT id, name, password FROM employee")
		if err != nil {
			fmt.Println("Error retrieving data from database:", err)
			return
		}

		message.SendData(w, employees)

		return
	}
	http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
}
