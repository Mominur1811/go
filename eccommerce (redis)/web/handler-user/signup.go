package handleruser

import (
	"bytes"
	"context"
	"crypto/sha1"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/smtp"
	"time"
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

	sendConfirmationEmail(insertedUser)

	utils.SendData(w, "", insertedUser)

}

func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func sendConfirmationEmail(user *db.User) error {
	from := "sourovsourov88@gmail.com"
	password := "rwxahqonrotbaozq"

	to := []string{
		user.Email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	t, err := template.ParseFiles("/home/mominur/ecommerce/html/signup-email.html")
	if err != nil {
		slog.Error("Failed to validate new user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		return err
	}

	otp := EncodeToString(6)
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Confirm SignUP \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Otp string
	}{
		Otp: otp,
	})
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		slog.Error("Failed to validate new user data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
		return err
	}

	err = db.GetRedisClient().Set(context.Background(), user.Email, otp, 5*time.Minute).Err()
	if err != nil {
		slog.Error("Failed to save otp in redis", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": user,
		}))
	}

	return err

}
