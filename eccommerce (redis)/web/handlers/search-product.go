package handlers

import (
	"ecommerce/db"
	"ecommerce/logger"
	"ecommerce/web/utils"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

func SearchProduct(w http.ResponseWriter, r *http.Request) {

	productParams, err := GetQueryParams(r.URL.Query())
	if err != nil {
		slog.Error("Failed to load query param", logger.Extra(map[string]any{
			"error":   err.Error(),
			"payload": productParams,
		}))
		utils.SendError(w, http.StatusPreconditionFailed, err.Error())
		return
	}

	//Go routine to find count of search
	countChan := make(chan db.CountResult)
	go db.GetProductRepo().GetCountProduct(productParams, countChan)

	productList, err := db.GetProductRepo().GetSearchResultProduct(productParams)
	if err != nil {
		utils.SendError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	result := <-countChan
	if result.Err != nil {
		utils.SendError(w, http.StatusExpectationFailed, result.Err.Error())
		return
	}

	//Send Product and Info as Json
	utils.SendData(w, map[string]interface{}{
		"Total Result":  result.Count,
		"Page No":       productParams.Page,
		"Total Page":    (result.Count + productParams.Limit - 1) / productParams.Limit,
		"Page Contents": len(productList)}, productList)
}

func GetQueryParams(queryParams url.Values) (db.ProductQueryParam, error) {

	//Get Price
	var productPrice int
	var err error
	if productPriceStr := queryParams.Get("product_price"); productPriceStr != "" {
		productPrice, err = strconv.Atoi(productPriceStr)
		if err != nil {
			return db.ProductQueryParam{}, err
		}
	}

	//Get page
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		return db.ProductQueryParam{}, err
	}

	//Get Page limit
	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		return db.ProductQueryParam{}, err
	}
	limit = min(25, max(limit, 3))

	return db.ProductQueryParam{
		ProductName:     queryParams.Get("product_name"),
		ProductCategory: queryParams.Get("product_category"),
		ProductPrice:    productPrice,
		Page:            page,
		Limit:           limit,
		SortType:        queryParams.Get("sort_type"),
	}, nil
}
