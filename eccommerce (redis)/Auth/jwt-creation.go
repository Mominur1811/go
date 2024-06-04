package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/db"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func GetAccessToken(userLogin db.Login, duration int) (string, error) {

	jwtData := map[string]interface{}{
		"expiration_time": time.Now().Add(time.Minute * time.Duration(duration)).Unix(),
		"username":        userLogin.Email,
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

	secretKey := config.GetConfig().JwtSecretKey
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write(payload)
	signature := hash.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(signature)

}
