package handleruser

import (
	auth "ecommerce/Auth"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func RefreshAccessToken(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	claims, err := GetPayloadData(token)
	if err != nil {
		slog.Error("Failed to claims data from jwt_token", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": claims,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, err := auth.GetAccessToken(claims, 2)
	if err != nil {
		slog.Error("Failed to generate jwt_token", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": newAccessToken,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendData(w, http.StatusOK, newAccessToken)

}

func GetPayloadData(str string) (db.Login, error) {

	token := strings.Split(str, ".")

	decodedPayload, err := base64.StdEncoding.DecodeString(token[1])
	if err != nil {
		return db.Login{}, err
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(decodedPayload, &claims); err != nil {

		return db.Login{}, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return db.Login{}, fmt.Errorf("unable to extract data from access token")
	}

	password, ok := claims["password"].(string)
	if !ok {
		return db.Login{}, fmt.Errorf("unable to extract password from access token")
	}

	return db.Login{Username: username, Password: password}, nil
}
