
package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func VenderDataRoutes(incomingRoutes *gin.Engine){
	routes := incomingRoutes.Group("/api")
	{
	routes.GET("/invender", controller.GetIndvisualVenders())
	routes.GET("/invender/:vender_id", controller.GetIndvisualVender())
	routes.POST("/invender/:vender_type", controller.CreateIndvisualVender())
	routes.PATCH("/invender/:vender_id", controller.UpdateIndvisualVender())
	}
}