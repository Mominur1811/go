package message

import (
	"fmt"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}) {
	fmt.Println(data)
	Send_Json(w, http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}
