package routes

import (
	Controllers "product_services/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp_product := r.Group("/api/v1/product")
	{
		grp_product.POST("create", Controllers.CreateProduct)
		grp_product.POST("create/:merchant_id", Controllers.CreateProductAPI)
		grp_product.GET("category/create/:name", Controllers.CreateCategory)
		grp_product.GET("subcategory/create/:category_id/:name", Controllers.CreateSubCategory)

		grp_product.POST("/update/:id", Controllers.UpdateProduct)

		grp_product.GET("business/:business_id", Controllers.GetProductByMerchantId)

		grp_product.GET(":product_id", Controllers.GetProductById)

		grp_product.GET("/search/:category_id/:name", Controllers.SearchProductByName)

		grp_product.GET("/best_seller", Controllers.GetBestSellerProduct)

		grp_product.GET("/active_state/:id", Controllers.ChangeProductActiveState)
	}

	return r
}
