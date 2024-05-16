package db

import (
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
