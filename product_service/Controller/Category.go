package controller

import (
	"net/http"
	Model "product_services/Model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category Model.ProductCategory

	category.Name = c.Params.ByName("name")

	err := Model.CreateCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, category)
}

func CreateSubCategory(c *gin.Context) {
	var subCategory Model.ProductSubCategory
	var err error

	subCategory.Name = c.Params.ByName("name")
	subCategory.CategoryId, err = strconv.Atoi(c.Params.ByName("category_id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = Model.CreateSubCategory(&subCategory)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, subCategory)
}
