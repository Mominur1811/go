package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func GetVerifyJwt(token map[string]string) (map[string]interface{}, error) {
	parts := strings.Split(token["token"], ".")
	if len(parts) != 3 {

		return nil, fmt.Errorf("invalid token formal")
	}

	if err := GetVerifyHeader(parts[0]); err != nil {
		return nil, err
	}

	if err := GetVerifySignature(parts[1], parts[2]); err != nil {
		return nil, err
	}

	payload, err := GetPayload(parts[1])
	if err != nil {
		return nil, err
	}

	return payload, err
}

func GetVerifyHeader(encodedHeader string) error {

	header, err := base64.RawURLEncoding.DecodeString(encodedHeader)
	if err != nil {
		return fmt.Errorf("error decoding header: %v", err)
	}

	var headerData map[string]interface{}
	if err := json.Unmarshal(header, &headerData); err != nil {
		return fmt.Errorf("error parsing header: %v", err)
	}

	alg, ok := headerData["alg"].(string)
	if !ok || alg != "HS256" {
		return fmt.Errorf("unsupported algorithm")
	}

	return nil
}

func GetVerifySignature(encodedPayload string, signature string) error {

	payload, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		return err
	}

	expectedSignature := GenerateSignature(payload)

	if signature == expectedSignature {
		return nil
	}
	return fmt.Errorf("signature verification failed")

}

func GetPayload(token string) (map[string]interface{}, error) {

	decodedPayload, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(decodedPayload, &claims); err != nil {

		return nil, err
	}

	expirationTime, ok := claims["expiration_time"].(float64)
	if !ok {
		return nil, fmt.Errorf("expiration_time field is not a float64")
	}

	if int64(expirationTime) > time.Now().Unix() {
		return claims, nil
	}

	return nil, fmt.Errorf("user is unauthorized")
}
