package db

import (
	"ecommerce/logger"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

// Product struct represents the product table
type Product struct {
	ProductName     string  `db:"product_name"       validate:"required,min=4,max=12"              json:"product_name"`
	ProductCategory string  `db:"product_category"   validate:"required,min=4,max=12"              json:"product_category"`
	ProductPrice    float64 `db:"product_price" validate:"required,checkType=float64,min=4,max=12" json:"product_price"`
	ProductQuantity int     `db:"product_quantity" validate:"required,checkType=int,min=4,max=12"  json:"product_quantity"`
}

// Use to take query params for product searching, pagination,sorting
type ProductQueryParam struct {
	ProductName     string
	ProductCategory string
	ProductPrice    int
	Page            int
	Limit           int
	SortType        string
}

// Use as channel data type of a go routine
type CountResult struct {
	Count int
	Err   error
}

type ProductRepo struct {
	productTableName string
}

var productRepo *ProductRepo

func InitProductRepo() {
	productRepo = &ProductRepo{productTableName: "product"}
}

func GetProductRepo() *ProductRepo {
	return productRepo
}

func (r *ProductRepo) GetSearchResultProduct(productParam ProductQueryParam) ([]*Product, error) {

	query := GetQueryBuilder().
		Select("product_name, product_category, product_price, product_quantity").
		From(r.productTableName)

	if productParam.ProductName != "" {
		query = query.Where(sq.Eq{"product_name": productParam.ProductName})
	}

	if productParam.ProductCategory != "" {
		query = query.Where(sq.Eq{"product_category": productParam.ProductCategory})
	}

	if productParam.ProductPrice > 0 {
		query = query.Where(sq.GtOrEq{"product_price": productParam.ProductPrice})
	}

	if productParam.SortType != "" {
		query = query.OrderBy("product_price " + productParam.SortType)
	}

	offset := ((productParam.Page) - 1) * productParam.Limit
	query = query.Limit(uint64(productParam.Limit)).Offset(uint64(offset))

	qry, args, err := query.ToSql()
	if err != nil {
		slog.Error(
			"Failed to create product result query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
		return nil, err
	}

	productList := []*Product{}
	err = GetReadDB().Select(&productList, qry, args...)
	if err != nil {
		slog.Error(
			"Failed to create product result query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
		return nil, err
	}

	return productList, nil
}

func (r *ProductRepo) GetCountProduct(productParam ProductQueryParam, coutChan chan CountResult) {

	query := GetQueryBuilder().
		Select("Count(*)").
		From(r.productTableName)

	if productParam.ProductName != "" {
		query = query.Where(sq.Eq{"product_name": productParam.ProductName})
	}

	if productParam.ProductCategory != "" {
		query = query.Where(sq.Eq{"product_category": productParam.ProductCategory})
	}

	if productParam.ProductPrice > 0 {
		query = query.Where(sq.GtOrEq{"product_price": productParam.ProductPrice})
	}

	qry, args, err := query.ToSql()
	if err != nil {
		slog.Error(
			"Failed to create product count query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)

	}

	var countProduct int
	err = GetReadDB().Get(&countProduct, qry, args...)
	if err != nil {
		slog.Error(
			"Failed product count query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
	}

	coutChan <- CountResult{Count: countProduct, Err: err}
}

func (r *ProductRepo) AddProduct(product Product) (*Product, error) {

	column := map[string]interface{}{
		"product_name":     product.ProductName,
		"product_category": product.ProductCategory,
		"product_price":    product.ProductPrice,
		"product_quantity": product.ProductQuantity,
	}
	var columns []string
	var values []any
	for columnName, columnValue := range column {

		columns = append(columns, columnName)
		values = append(values, columnValue)

	}
	qry, args, err := GetQueryBuilder().
		Insert(r.productTableName).
		Columns(columns...).
		Suffix(`
			RETURNING 
			product_name,
			product_category,
			product_price,
			product_quantity
		`).
		Values(values...).
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create product insert query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
		return nil, err
	}
	// Execute the SQL query and get the result
	var insertedProduct Product
	err = GetReadDB().QueryRow(qry, args...).Scan(&insertedProduct.ProductName, &insertedProduct.ProductCategory, &insertedProduct.ProductPrice, &insertedProduct.ProductQuantity)
	if err != nil {
		slog.Error(
			"Failed to execute product insert query",
			logger.Extra(map[string]interface{}{
				"error": err.Error(),
				"query": qry,
				"args":  args,
			}),
		)
	}

	return &insertedProduct, nil
}
