package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func ProductDetailRouter(incomingRoutes *gin.Engine){
	routes := incomingRoutes.Group("/api")
	{
	routes.GET("/product_detail", controller.GetProductDetails())
	routes.GET("/product_detail/:vender_id", controller.GetProductDetail())
	routes.POST("/product_detail", controller.CreateProductDetail())
	routes.PATCH("/product_detail/:vender_id", controller.UpdateProductDetail())
	}
}