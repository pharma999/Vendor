
package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func VenderDataRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invender", controller.GetIndvisualVenders())
	incomingRoutes.GET("/invender/:vender_id", controller.GetIndvisualVender())
	incomingRoutes.POST("/invender/:vender_type", controller.CreateIndvisualVender())
	incomingRoutes.PATCH("/invender/:vender_id", controller.UpdateIndvisualVender())
}