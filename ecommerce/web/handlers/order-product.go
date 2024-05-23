package handlers

import (
	"ecommerce/db"
	"ecommerce/web/messages"
	"encoding/json"
	"net/http"
)

func NewOrder(w http.ResponseWriter, r *http.Request) {

	newOrder := db.Order{}

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	if err := db.AddOrder(newOrder); err != nil {
		messages.SendError(w, http.StatusFailedDependency, err.Error(), "")
		return
	}

	messages.SendData(w, "", newOrder)
}
