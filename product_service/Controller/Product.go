package controller

import (
	"fmt"
	"net/http"
	Model "product_services/Model"
	"product_services/Utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetProductByMerchantId(c *gin.Context) {
	var products []Model.Product

	err := Model.GetProductByMerchantId(&products, c.Params.ByName("business_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	var product Model.Product

	err := Model.GetProductById(&product, c.Params.ByName("product_id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product Model.ProductRequest

	c.BindJSON(&product)

	photosUrl := make([]string, len(product.Photos))

	Utils.LoadEnv()
	sess := Utils.ConnectAws()

	for index, item := range product.Photos {
		photosUrl[index] = Utils.RandStringRunes(20)
		go Utils.AddFileToS3(sess, item, photosUrl[index])
	}

	photosString := strings.Join(photosUrl, ",")

	// Creating the actual product

	var productCreate Model.Product
	fmt.Println(product.Price)
	productCreate.Photos = photosString
	productCreate.Price = product.Price
	productCreate.Active = true
	productCreate.Halal = true
	productCreate.Name = product.Name
	productCreate.Description = product.Description
	productCreate.ReviewsCount = product.ReviewsCount
	productCreate.Score = product.Score
	productCreate.Sold = product.Sold
	productCreate.CategoryId = product.Id
	productCreate.MerchantId = 46

	err := Model.CreateProduct(&productCreate)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, productCreate)
}

func CreateProductAPI(c *gin.Context) {
	var product Model.ProductRequest
	merchantId := c.Params.ByName("merchant_id")

	c.BindJSON(&product)
	fmt.Println(product.CategoryId)
	photosUrl := make([]string, len(product.Photos))

	Utils.LoadEnv()
	sess := Utils.ConnectAws()

	for index, item := range product.Photos {
		photosUrl[index] = Utils.RandStringRunes(20)
		go Utils.AddFileToS3(sess, item, photosUrl[index])
	}

	photosString := strings.Join(photosUrl, ",")

	// Creating the actual product
	var err error

	var productCreate Model.Product
	productCreate.Photos = photosString
	productCreate.Price = product.Price
	productCreate.Active = true
	productCreate.Halal = true
	productCreate.Name = product.Name
	productCreate.Description = product.Description
	productCreate.ReviewsCount = product.ReviewsCount
	productCreate.Score = product.Score
	productCreate.Sold = product.Sold
	productCreate.CategoryId = product.CategoryId
	productCreate.MerchantId, err = strconv.Atoi(merchantId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = Model.CreateProduct(&productCreate)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, productCreate)
}

func SearchProductByName(c *gin.Context) {
	var products []Model.Product

	err := Model.GetProductByName(&products, c.Params.ByName("name"), c.Params.ByName("category_id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetBestSellerProduct(c *gin.Context) {
	var products []Model.Product

	err := Model.GetBestSellerProduct(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	var product Model.Product
	err := Model.GetProductById(&product, c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var incomingProduct Model.Product
	c.BindJSON(&incomingProduct)

	product.Name = incomingProduct.Name
	product.Description = incomingProduct.Description
	product.Price = incomingProduct.Price

	err = Model.UpdateProduct(&product)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, product)

}

func ChangeProductActiveState(c *gin.Context) {
	var product Model.Product

	err := Model.GetProductById(&product, c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if product.Active {
		product.Active = false
	} else {
		product.Active = true
	}

	err = Model.UpdateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}
