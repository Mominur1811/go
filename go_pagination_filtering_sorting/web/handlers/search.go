package handlers

import (
	"net/http"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web/messages"
)

func Search(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	pageNo := queryParams.Get("page")
	limit := queryParams.Get("limit")
	orderKey := queryParams.Get("orderKey")
	orderType := queryParams.Get("sort")
	colName := queryParams.Get("colName")
	colValue := queryParams.Get("colValue")

	val, err := db.MakeQuery(pageNo, limit, orderKey, orderType, colName, colValue)

	if err != nil {

		messages.SendError(w, http.StatusExpectationFailed, err.Error(), "")
		return
	}

	messages.SendData(w, val)

}
