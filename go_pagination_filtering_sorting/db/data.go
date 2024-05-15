package db

import (
	"errors"
	"fmt"
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

func GetInt(val1 string, val2 *int) error {

	v, err := strconv.Atoi(val1)
	*val2 = v
	return err
}

func GetName(val1 string) (string, error) {
	switch val1 {

	case "pname":
		return "pname", nil
	case "cname":
		return "customer", nil
	case "date":
		return "orderdate", nil
	case "price":
		return "price", nil
	default:
		return "", errors.New("unknown search parameter")

	}
}

func MakeQuery(pageNo string, limit string, orderkey string, orderType string, colName string, colValue string) ([]Product, error) {

	newPage := 1
	newLimit := 2
	var product []Product

	var srtCriteria, matchName string
	if orderkey != "" {
		srtCriteria = " ORDER BY " + orderkey + " " + orderType + " "
	}

	if err := GetInt(pageNo, &newPage); err != nil {
		return product, err
	}

	if err := GetInt(limit, &newLimit); err != nil {
		return product, err
	}

	if colName != "" {
		col, err := GetName(colName)
		if err != nil {
			return product, err
		}

		matchName = fmt.Sprintf(" WHERE %s = '%s' ", col, colValue)
	}
	offset := (newPage - 1) * newLimit
	str := fmt.Sprintf("SELECT pname, price, orderdate, customer FROM products %s%s  LIMIT %d OFFSET %d", matchName, srtCriteria, newLimit, offset)
	fmt.Println(str)
	db := GetDB()
	err := db.Select(&product, str)
	return product, err

}
