package db

import "fmt"

func BuildSearchQuery(queryParam ProductQueryParam) (string, string) {

	queryWithPaginationFilter := `SELECT product_name, product_category, product_price, product_quantity  FROM product WHERE 1 = 1`
	searchResultCount := `SELECT COUNT(product_id)  FROM product WHERE 1 = 1`
	whereClause := ""

	if queryParam.ProductName != "" {
		whereClause += fmt.Sprintf(" AND product_name = '%s'", queryParam.ProductName)
	}

	if queryParam.ProductCategory != "" {
		whereClause += fmt.Sprintf(" AND product_category = '%s'", queryParam.ProductCategory)
	}

	if queryParam.ProductPrice != 0 {
		whereClause += fmt.Sprintf(" AND product_price <= %d", queryParam.ProductPrice)
	}

	//Merge Where condition
	searchResultCount += whereClause
	queryWithPaginationFilter += whereClause

	//Merge Sorting order
	queryWithPaginationFilter += fmt.Sprintf(" ORDER BY product_price %s", queryParam.SortType)

	//Merge Pagination clause
	offset := (queryParam.Page - 1) * queryParam.Limit
	queryWithPaginationFilter += fmt.Sprintf(" LIMIT %d OFFSET %d", queryParam.Limit, offset)
	return queryWithPaginationFilter, searchResultCount

}
