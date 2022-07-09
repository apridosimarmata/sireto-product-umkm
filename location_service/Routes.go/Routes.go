package Routes

import (
	"location_service/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp_province := r.Group("/api/v1/province")
	{
		grp_province.POST("create", Controllers.CreateProvince)
		grp_province.GET("all", Controllers.GetAllProvinces)
	}

	grp_city := r.Group("/api/v1/city")
	{
		grp_city.POST("create", Controllers.CreateCity)
		grp_city.GET("/all/:province_id", Controllers.GetCitiesByProvinceId)
	}

	return r
}
