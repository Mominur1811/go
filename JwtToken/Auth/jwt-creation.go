package auth

import (
	"JwtToken/db"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func GetJwtToken(userLogin db.LoginData) (string, error) {

	jwtData := map[string]interface{}{
		"expiration_time": time.Now().Add(time.Minute * 1).Unix(),
		"username":        userLogin.UserName,
		"password":        userLogin.Password,
	}

	payload, err := json.Marshal(jwtData)
	if err != nil {
		return "", err
	}

	encodedPayload := base64.StdEncoding.EncodeToString(payload)

	//Generate Header part
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

	// Concatenate header, payload and signature
	jwt := fmt.Sprintf("%s.%s.%s", header, encodedPayload, GenerateSignature(payload))
	return jwt, nil

}

func GenerateSignature(payload []byte) string {

	secretKey := []byte("your_secret_key")
	hash := hmac.New(sha256.New, secretKey)
	hash.Write(payload)
	signature := hash.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(signature)

}
