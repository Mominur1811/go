package db

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

func AddProduct(newProduct Product) error {

	db := GetDB()
	_, err := db.Exec(`INSERT INTO product (product_name, product_category, product_price, product_quantity)
	                  VALUES ($1, $2, $3, $4)`,
		newProduct.ProductName, newProduct.ProductCategory, newProduct.ProductPrice, newProduct.ProductQuantity)

	return err
}

func GetProductList(str string) ([]Product, error) {

	db := GetDB()
	var productList []Product
	err := db.Select(&productList, str)
	return productList, err
}

func GetSearchCount(str string, countChan chan CountResult) {
	db := GetDB()
	var count int
	err := db.Get(&count, str)
	result := CountResult{Count: count, Err: err}
	countChan <- result
}
