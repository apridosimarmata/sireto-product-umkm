package Routes

import (
	"umkm_service/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp_umkm := r.Group("/api/v1/umkm")
	{
		grp_umkm.POST("create", Controllers.CreateUMKM)
		grp_umkm.POST("create/:user_id", Controllers.CreateUMKMApi)

		grp_umkm.POST("category", Controllers.CreateUMKMCategory)
		grp_umkm.GET(":id", Controllers.GetUMKMById)
		grp_umkm.GET("categories", Controllers.GetAllUMKMCategories)

		grp_umkm.POST("subcategory", Controllers.CreateUMKMSubCategory)
		grp_umkm.GET("subcategories/:category_id", Controllers.GetUMKMSubCategoriesByCategoryId)

		grp_umkm.GET("location/:location_id", Controllers.GetUMKMByLocationId)

		grp_umkm.PATCH("score", Controllers.AddUMKMScore)

		grp_umkm.GET("user/:owner_id", Controllers.GetUMKMSByOwnerId)

		grp_umkm.GET("favorite", Controllers.GetFavoriteUMKM)
		grp_umkm.GET("search/:name", Controllers.SearchUMKMByName)

		grp_umkm.GET("open/:id", Controllers.OpenBusiness)
		grp_umkm.GET("close/:id", Controllers.CloseBusiness)

	}

	return r
}
