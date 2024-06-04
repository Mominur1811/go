package handleruser

import (
	"ecommerce/config"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	//Search User Exist or Not
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
	jwtToken, err := createToken(userLogin.Email, 5)
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
	refreshToken, err := createToken(userLogin.Email, 50)
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

// JWT CREATION
func createToken(username string, lifeTime int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * time.Duration(lifeTime)).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.GetConfig().JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
