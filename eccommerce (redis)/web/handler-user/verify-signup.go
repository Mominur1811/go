package handleruser

import (
	"context"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

type OTP struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

func VerifySignUp(w http.ResponseWriter, r *http.Request) {

	var otp OTP
	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		slog.Error("Failed to decode user login data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": otp,
		}))

		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	actualOtp := db.GetRedisClient().Get(context.Background(), otp.Email)

	if actualOtp.Val() != otp.Otp {

		utils.SendData(w, "", "Invalid Otp")
		return
	}

	err := db.GetUserLoginRepo().UpdateCheckMarkUser(otp.Email)
	if err != nil {
		slog.Error("Failed to Update is_valid column of customer", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": otp,
		}))
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
	}
	utils.SendData(w, "", "Successfullty verified user")
}
