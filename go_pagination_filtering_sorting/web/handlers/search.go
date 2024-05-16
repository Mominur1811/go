package handlers

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"example.com/go_pagination_filtering_sorting/db"
	"example.com/go_pagination_filtering_sorting/web/messages"
)

func Search(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	pageNum, limit, err := GetPageInfoCount(queryParams.Get("page"), queryParams.Get("limit"))
	if err != nil {
		messages.SendError(w, http.StatusPreconditionFailed, err.Error(), "")
	}

	//Load Query String
	queryString, infoString, err := GetQuery(queryParams, pageNum, limit)
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

	count, err := concurrency(infoString)
	if err != nil {
		messages.SendError(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	fmt.Println(count)
	additionalInfo := getAdditionalInfo(count, len(product), pageNum, limit)
	messages.SendData(w, additionalInfo, product)

}

func concurrency(infoString string) (int, error) {
	var count int
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go routine1(infoString, &count, &err, &wg)

	wg.Wait()
	return count, err
}

func routine1(infoString string, count *int, err *error, wg *sync.WaitGroup) {
	defer wg.Done()

	*count, *err = db.QueryPaginationFilterInfo(infoString)
}

// Give Additional Info To The Fontend

func getAdditionalInfo(count int, dataCount int, pageNum int, limit int) interface{} {

	totalPage := math.Ceil(float64(count) / float64(limit))

	return map[string]interface{}{
		"Total Page":     totalPage,
		"Total Contents": count,
		"Page No":        pageNum,
		"Page Contents":  dataCount,
	}

}

// Generate Query String
func GetQuery(queryParams url.Values, pageNum int, limit int) (string, string, error) {

	sortKey := queryParams.Get("sortKey")
	sortType := queryParams.Get("sortOrder")
	searchString := queryParams.Get("search")

	pagination := GetPagination(pageNum, limit)

	filter, err := GetFilter(sortKey, sortType)
	if err != nil {
		return "", "", err
	}

	search := GetSearch(searchString)

	offsetQuery := fmt.Sprintf("SELECT pname, price, orderdate, customer FROM products %s %s %s", search, filter, pagination)
	countQuery := fmt.Sprintf("SELECT COUNT(pname) FROM products %s", search)
	return offsetQuery, countQuery, nil
}

// Get Page No and Page Contents Limit
func GetPageInfoCount(pageNum string, contentLimit string) (int, int, error) {
	page, err := strconv.Atoi(pageNum)
	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(contentLimit)
	if err != nil {
		return 0, 0, err
	}

	//Set Valid Page No
	page = max(page, 1)

	//Set Valid Page Contents Limit
	limit = max(limit, 2)
	limit = min(limit, 25)

	return page, limit, nil
}

// Generate Pagination String Like Limit * OFFSET *
func GetPagination(page int, limit int) string {

	//Calc Offset
	offset := (page - 1) * limit
	pagination := fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	return pagination

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
