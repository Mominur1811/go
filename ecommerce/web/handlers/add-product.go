package handlers

import (
	"ecommerce/db"
	"ecommerce/web/messages"
	"encoding/json"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct db.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {

		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	if err := db.AddProduct(newProduct); err != nil {
		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	messages.SendData(w, "", newProduct)
}
