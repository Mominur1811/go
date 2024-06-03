package handlers

import (
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var newUser db.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		slog.Error("Failed to decode new user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newUser,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error())
		return
	}

	if err := utils.ValidateStruct(newUser); err != nil {
		slog.Error("Failed to validate new user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newUser,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}
	userRepo := db.GetUserRegisterRepo()
	var user *db.User
	var err error
	if user, err = userRepo.InsertNewUser(&newUser); err != nil {
		slog.Error("Failed to insert new user data in db", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newUser,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.SendData(w, "", user)

}
