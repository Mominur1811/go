package db

import "time"

// Order struct represents the order table
type Order struct {
	Username      string    `db:"username"        json:"username"`
	ProductID     int       `db:"product_id"      json:"product_id"`
	OrderDate     time.Time `db:"order_date"      json:"order_date"`
	OrderQuantity int       `db:"order_quantity"  json:"order_quantity"`
	TotalPay      float64   `db:"total_pay"       json:"total_pay" `
}

func AddOrder(newOrder Order) error {

	db := GetDB()
	_, err := db.Exec(`INSERT INTO "order" (username, product_id, order_date, order_quantity, total_pay) 
                      VALUES ($1, $2, $3, $4, $5)`,
		newOrder.Username, newOrder.ProductID, newOrder.OrderDate, newOrder.OrderQuantity, newOrder.TotalPay)
	return err
}
