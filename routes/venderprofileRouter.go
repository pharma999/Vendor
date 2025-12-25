package routes

import ( 
	"github.com/gin-gonic/gin"
	controller "github.com/pharma999/vender/controller"
)

func VenderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/vender", controller.GetVenders())
	incomingRoutes.GET("/vender/:vender_id", controller.GetVender())
	incomingRoutes.POST("/vender", controller.CreateVender())
	incomingRoutes.PATCH("/vender/:vender_id", controller.UpdateVender())
}