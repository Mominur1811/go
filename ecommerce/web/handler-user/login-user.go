package handleruser

import (
	auth "ecommerce/Auth"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLogin db.Login
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		slog.Error("Failed to decode user login data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": userLogin,
		}))

		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	if err := utils.ValidateStruct(userLogin); err != nil {
		slog.Error("Failed to validate user login data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": userLogin,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	//User Login Validation
	userLogin.Password = hashPassword(userLogin.Password)
	userRepo := db.GetUserLoginRepo()
	if _, err := userRepo.FindUser(&userLogin); err != nil {
		slog.Error("Failed to login", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": userLogin,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	//Get Jwt Token
	jwtToken, err := auth.GetAccessToken(userLogin, 5)
	if err != nil {
		slog.Error("Failed to get access token", logger.Extra(map[string]any{
			"error":     err.Error(),
			"payload":   userLogin,
			"jwt_token": jwtToken,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	//Get Refresh Token
	refreshToken, err := auth.GetAccessToken(userLogin, 5)
	if err != nil {
		slog.Error("Failed to get refresh token", logger.Extra(map[string]any{
			"error":         err.Error(),
			"payload":       userLogin,
			"refresh_token": jwtToken,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	utils.SendData(w, http.StatusOK, map[string]interface{}{
		"Access Token ": jwtToken,
		"Refresh Token": refreshToken,
	})
}
