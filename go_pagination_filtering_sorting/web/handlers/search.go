package handlers

import (
	"net/http"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web/messages"
)

func Search(w http.ResponseWriter, r *http.Request) {

	//Generate query string
	var queryString string
	if err := db.GetQueryString(r, &queryString); err != nil {
		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	//load data from database
	var product []db.Product
	if err := db.PaginationFilterSearch(queryString, &product); err != nil {
		messages.SendError(w, http.StatusAccepted, err.Error(), "")
		return
	}
	messages.SendData(w, product)

}
