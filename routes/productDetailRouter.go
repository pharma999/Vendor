package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func ProductDetailRouter(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/product_detail", controller.GetProductDetails())
	incomingRoutes.GET("/product_detail/:vender_id", controller.GetProductDetail())
	incomingRoutes.POST("/product_detail", controller.CreateProductDetail())
	incomingRoutes.PATCH("/product_detail/:vender_id", controller.UpdateProductDetail())
}