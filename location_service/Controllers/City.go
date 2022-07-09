package Controllers

import (
	"location_service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCity(c *gin.Context) {
	var city Models.City
	c.BindJSON(&city)

	err := Models.CreateCity(&city)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func GetCitiesByProvinceId(c *gin.Context) {
	province_id := c.Params.ByName("province_id")
	var cities []Models.City

	err := Models.GetCitiesByProvinceId(&cities, province_id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, cities)
}
