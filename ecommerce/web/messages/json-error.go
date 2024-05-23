package messages

import "net/http"

func SendError(w http.ResponseWriter, status int, message string, data interface{}) {

	SendJson(w, status, map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
