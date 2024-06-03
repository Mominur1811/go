package handlers

import (
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct db.Product
	var err error

	if err = json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		slog.Error("Failed to decode product data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newProduct,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error())
		return
	}

	if err = utils.ValidateStruct(newProduct); err != nil {
		slog.Error("Failed to validate product data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newProduct,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error())
		return
	}

	var insertedProduct *db.Product
	if insertedProduct, err = db.GetProductRepo().AddProduct(newProduct); err != nil {
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.SendData(w, "", insertedProduct)
}
