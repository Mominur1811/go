package message

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Send_Json(w http.ResponseWriter, status int, data interface{}) {

	w.Header().Add("Content-Type", "application/json")

	str, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Infinite loop")
		message := "Convertion to Json failed"
		Send_Error(w, status, message, data)
		return
	}
	w.WriteHeader(status)
	w.Write(str)

}
