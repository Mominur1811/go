package handleruser

import (
	"bytes"
	"crypto/rand"
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"net/smtp"
)

func ResetPass(w http.ResponseWriter, r *http.Request) {

	var resetPass db.Login
	var err error
	if err = json.NewDecoder(r.Body).Decode(&resetPass); err != nil {
		slog.Error("Failed to load json", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": resetPass,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.ValidateStruct(resetPass); err != nil {
		slog.Error("Error in json data", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": resetPass,
		}))
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	var email string
	resetPass.Password = hashPassword(resetPass.Password)
	if email, err = db.GetUserLoginRepo().FindUser(&resetPass); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOtp(email)

}

func sendOtp(email string) error {

	from := "sourovsourov88@gmail.com"
	password := "rwxahqonrotbaozq"

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	t, err := template.ParseFiles("/home/mominur/ecommerce/web/handler-user/reset-pass.html")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	otp := EncodeToString(6)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Otp string
	}{
		Otp: otp,
	})
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//db.GetUserLoginRepo().InsertOtp(email, strconv.Itoa(opt))
	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	return err
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
