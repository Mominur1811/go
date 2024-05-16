package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web/messages"
)

func Insert(w http.ResponseWriter, r *http.Request) {

	var newEntry db.Product

	if err := json.NewDecoder(r.Body).Decode(&newEntry); err != nil {

		messages.SendError(w, http.StatusBadRequest, "Failed to load json", "")
		return
	}

	if err := db.InsertInTable(newEntry); err != nil {

		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	messages.SendJson(w, http.StatusCreated, newEntry)

}
