package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func VenderRoutes(incomingRoutes *gin.Engine){
	routes := incomingRoutes.Group("/api")
	{
	routes.GET("/vender", controller.GetVenders())
	routes.GET("/vender/:vender_id", controller.GetVender())
	routes.POST("/vender", controller.CreateVender())
	routes.PATCH("/vender/:vender_id", controller.UpdateVender())
	}
}