package db

import (
	"time"
)

func InsertInTable(info Product) error {

	db := GetDB()
	info.OrderDate = time.Now()
	_, err := db.NamedExec("INSERT INTO products(pname, price, orderdate, customer) VALUES(:pname, :price, :orderdate, :customer)", info)
	return err
}

// Get data after applying pagination search and filtering
func QueryPaginationFilter(queryString string) ([]Product, error) {

	db := GetDB()
	var product []Product
	err := db.Select(&product, queryString)
	return product, err
}

func QueryPaginationFilterInfo(queryInfo string) (int, error) {

	db := GetDB()
	var count int
	err := db.Get(&count, queryInfo)
	return count, err
}
