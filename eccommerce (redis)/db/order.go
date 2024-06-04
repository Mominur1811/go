package db

import (
	"ecommerce/logger"
	"log/slog"
	"time"
)

// Order struct represents the order table
type Order struct {
	Username      string    `db:"username"        json:"username"`
	ProductID     int       `db:"product_id"      json:"product_id"`
	OrderDate     time.Time `db:"order_date"      json:"order_date"`
	OrderQuantity int       `db:"order_quantity"  json:"order_quantity"`
	TotalPay      float64   `db:"total_pay"       json:"total_pay" `
}

type OrderRepo struct {
	orderTableName string
}

var orderRepo *OrderRepo

func InitOrderRepo() {

	orderRepo = &OrderRepo{orderTableName: `"order"`}
}

func GetOrderRepo() *OrderRepo {
	return orderRepo
}

func (r *OrderRepo) InsertNewOrder(newOrder *Order) (*Order, error) {
	db := GetWriteDB() // Assuming GetWriteDB() returns a *sql.DB for writing operations

	// Define the columns and values to be inserted
	column := map[string]interface{}{
		"username":       newOrder.Username,
		"product_id":     newOrder.ProductID,
		"order_date":     newOrder.OrderDate,
		"order_quantity": newOrder.OrderQuantity,
		"total_pay":      newOrder.TotalPay,
	}

	var columns []string
	var values []any
	for columnName, columnValue := range column {

		columns = append(columns, columnName)
		values = append(values, columnValue)

	}

	// Build the insert query using squirrel
	qry, args, err := GetQueryBuilder().
		Insert(r.orderTableName).
		Columns(columns...).
		Suffix(`
			RETURNING 
			product_id,
			order_date,
			order_quantity,
		    total_pay
		`).
		Values(values...).
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create order insert query",
			logger.Extra(map[string]interface{}{"error": err.Error()}),
		)
		return nil, err
	}

	// Execute the SQL query and get the result
	var insertedOrder Order
	err = db.QueryRow(qry, args...).Scan(
		&insertedOrder.Username,
		&insertedOrder.ProductID,
		&insertedOrder.OrderDate,
		&insertedOrder.OrderQuantity,
		&insertedOrder.TotalPay,
	)
	if err != nil {
		slog.Error(
			"Failed to execute order insert query",
			logger.Extra(map[string]interface{}{"error": err.Error()}),
		)
		return nil, err
	}

	return &insertedOrder, nil
}
