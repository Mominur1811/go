package handlers

import (
	"encoding/json"
	"fmt"
	"gobasic/model"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		var update_user User
		err := json.NewDecoder(r.Body).Decode(&update_user)

		if err != nil {

			http.Error(w, "Error user to update", http.StatusBadRequest)
			return
		}

		db := model.GetDB()
		query := `UPDATE employee SET name = $1, password = $2 WHERE id = $3`
		_, err = db.Exec(query, update_user.Name, update_user.Password, update_user.ID)
		if err != nil {
			fmt.Println("Error updating record:", err)
			return
		}

		fmt.Println("Record updated successfully")

		return

	}
}
