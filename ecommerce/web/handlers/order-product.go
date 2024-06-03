package handlers

import (
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func NewOrder(w http.ResponseWriter, r *http.Request) {

	newOrder := db.Order{}
	var err error

	if err = json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		slog.Error("Failed to decode new order data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newOrder,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error())
		return
	}

	if err = utils.ValidateStruct(newOrder); err != nil {
		slog.Error("Failed to validate new order data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newOrder,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	var insertedOrder *db.Order
	if insertedOrder, err = db.GetOrderRepo().InsertNewOrder(&newOrder); err != nil {
		slog.Error("Failed to add new order data in db", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newOrder,
		}))
		utils.SendError(w, http.StatusFailedDependency, err.Error())
		return
	}

	utils.SendData(w, "", insertedOrder)
}
