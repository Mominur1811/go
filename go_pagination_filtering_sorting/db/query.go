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

func PaginationFilterSearch(queryString string, product *[]Product) error {

	db := GetDB()
	err := db.Select(product, queryString)
	return err
}
