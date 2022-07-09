package Controllers

import (
	"location_service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProvince(c *gin.Context) {
	var province Models.Province
	c.BindJSON(&province)

	err := Models.CreateProvince(&province)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func GetAllProvinces(c *gin.Context) {
	var provinces []Models.Province

	err := Models.GetAllProvinces(&provinces)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, provinces)
}
