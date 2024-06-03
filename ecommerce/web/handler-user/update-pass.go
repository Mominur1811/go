package handleruser

import (
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"encoding/json"
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

	//opt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	//val := strconv.Itoa(opt)
	message := []byte("Your Secret key is : ")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	//fmt.Println(message, email, strconv.Itoa(opt))

	//db.GetUserLoginRepo().InsertOtp(email, strconv.Itoa(opt))
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	return err
}
