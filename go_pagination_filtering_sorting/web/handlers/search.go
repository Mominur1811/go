package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web/messages"
)

func Search(w http.ResponseWriter, r *http.Request) {

	//Load Query String
	queryString, err := GetQuery(r)
	if err != nil {
		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
		return
	}

	//Load Data From Database
	product, err := db.QueryPaginationFilter(queryString)
	if err != nil {
		messages.SendError(w, http.StatusAccepted, err.Error(), "")
		return
	}
	messages.SendData(w, product)

}

// Generate Query String
func GetQuery(r *http.Request) (string, error) {

	queryParams := r.URL.Query()
	pageNo := queryParams.Get("page")
	limit := queryParams.Get("limit")
	sortKey := queryParams.Get("sortKey")
	sortType := queryParams.Get("sortOrder")
	searchString := queryParams.Get("search")

	pagination, err := GetPagination(pageNo, limit)
	if err != nil {
		return "", err
	}

	filter, err := GetFilter(sortKey, sortType)
	if err != nil {
		return "", err
	}

	search := GetSearch(searchString)

	str := fmt.Sprintf("SELECT pname, price, orderdate, customer FROM products %s %s %s", search, filter, pagination)

	return str, nil
}

// Generate Pagination String Like Limit * OFFSET *
func GetPagination(pageNo string, pageLimit string) (string, error) {

	page, err := strconv.Atoi(pageNo)
	if err != nil {
		return "", err
	}

	limit, err := strconv.Atoi(pageLimit)
	if err != nil {
		return "", err
	}

	//Set Valid Page No
	page = max(page, 1)

	//Set Valid Page Contents Limit
	limit = max(limit, 10)
	limit = min(limit, 25)

	//Calc Offset
	offset := (page - 1) * limit

	pagination := fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	return pagination, nil

}

// Generate Filter Like ORDER BY * ASC / DESC
func GetFilter(filterKey string, sortType string) (string, error) {

	if sortType == "" {
		sortType = "ASC"
	}
	var filter string
	switch filterKey {

	case "pname":
		filter = fmt.Sprintf(" ORDER BY %s %s", filterKey, sortType)
	case "price":
		filter = fmt.Sprintf(" ORDER BY %s %s", filterKey, sortType)
	case "orderdate":
		filter = fmt.Sprintf(" ORDER BY %s %s", filterKey, sortType)
	case "customer":
		filter = fmt.Sprintf(" ORDER BY %s %s", filterKey, sortType)
	case "":
		filter = " ORDER BY price ASC"
	default:
		return "", errors.New("invalid column name")
	}

	return filter, nil

}

// Generate Search String Like WHERE * ILIKE *
func GetSearch(search string) string {
	return fmt.Sprintf(" WHERE pname ILIKE '%%%s%%' OR CAST(price AS TEXT) ILIKE '%%%s%%' OR TO_CHAR(orderdate, 'YYYY-MM-DD HH24:MI:SS') ILIKE '%%%s%%' OR customer ILIKE '%%%s%%'",
		search, search, search, search)
}
