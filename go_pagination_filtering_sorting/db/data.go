package db

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	Pname     string    `db:"pname"         json:"pname"`
	Price     uint      `db:"price"         json:"price"`
	OrderDate time.Time `db:"orderdate"     json:"orderdate"`
	Customer  string    `db:"customer"      json:"customer"`
}

const (
	host    = "localhost"
	port    = "5432"
	dbName  = "Product_Info"
	sslMode = "disable"
)

var (
	user     = "root"
	password = "admin"
)

func GetPaginationString(pageNo string, pageLimit string, pagination *string) error {

	page, err1 := strconv.Atoi(pageNo)

	limit, err2 := strconv.Atoi(pageLimit)

	if err1 != nil || err2 != nil {
		return errors.New("error conveting pageNo or pageLimit")
	}

	//Set Valid Page No
	page = max(page, 1)
	page = min(page, 10)

	//Set Valid Page Contents Limit
	limit = max(limit, 1)
	limit = min(limit, 100)

	offset := (page - 1) * limit

	*pagination = fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	return nil

}

func GetFilterString(sortKey string, sortType string, filter *string) error {

	if sortType == "" {
		sortType = "ASC"
	}
	switch sortKey {

	case "pname":
		*filter = fmt.Sprintf(" ORDER BY %s %s", sortKey, sortType)
	case "price":
		*filter = fmt.Sprintf(" ORDER BY %s %s", sortKey, sortType)
	case "orderdate":
		*filter = fmt.Sprintf(" ORDER BY %s %s", sortKey, sortType)
	case "customer":
		*filter = fmt.Sprintf(" ORDER BY %s %s", sortKey, sortType)
	case "":
		*filter = " ORDER BY price ASC"
	default:
		return errors.New("invalid column name")
	}

	return nil

}

func GetSearchString(search string) string {
	return fmt.Sprintf(" WHERE pname ILIKE '%%%s%%' OR CAST(price AS TEXT) ILIKE '%%%s%%' OR TO_CHAR(orderdate, 'YYYY-MM-DD HH24:MI:SS') ILIKE '%%%s%%' OR customer ILIKE '%%%s%%'", search, search, search, search)
}

func GetQueryString(r *http.Request, str *string) error {

	queryParams := r.URL.Query()
	pageNo := queryParams.Get("page")
	limit := queryParams.Get("limit")
	sortKey := queryParams.Get("sortKey")
	sortType := queryParams.Get("sortOrder")
	searchString := queryParams.Get("search")

	var pagination, filter string

	if err := GetPaginationString(pageNo, limit, &pagination); err != nil {
		return err
	}

	if err := GetFilterString(sortKey, sortType, &filter); err != nil {
		return err
	}

	search := GetSearchString(searchString)

	*str = fmt.Sprintf("SELECT pname, price, orderdate, customer FROM products %s %s %s", search, filter, pagination)
	fmt.Println(*str)

	return nil

}
