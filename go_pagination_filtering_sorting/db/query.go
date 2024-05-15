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
