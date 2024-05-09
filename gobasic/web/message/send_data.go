package message

import (
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}) {
	Send_Json(w, http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}
