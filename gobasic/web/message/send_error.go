package message

import "net/http"

func Send_Error(w http.ResponseWriter, status int, message string, data interface{}) {

	Send_Json(w, status, map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
