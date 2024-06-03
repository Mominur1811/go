package handleruser

import (
	"crypto/sha1"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

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

	newUser.Password = hashPassword(newUser.Password)

	var insertedUser *db.User
	var err error
	if insertedUser, err = db.GetUserRegisterRepo().InsertNewUser(&newUser); err != nil {
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.SendData(w, "", insertedUser)

}

func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
