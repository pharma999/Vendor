package main

import (

 "github.com/gin-gonic/gin"
 
 "github.com/pharma999/vender/routes"

 "github.com/pharma999/vender/config"
 _ "github.com/pharma999/vender/docs"
 ginSwagger "github.com/swaggo/gin-swagger"
 swaggerFiles "github.com/swaggo/files"
)

//var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "vender")

// @title Vender API
// @version 1.0
// @description API for managing venders
// @host localhost:8000
// @BasePath /api
func main() {
	
	config.LoadConfig()
	
	port := config.GetEnv("PORT")

	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.VenderRoutes(router)
	routes.VenderDataRoutes(router)
	routes.ProductDetailRouter(router)
	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	router.Run(":" + port)
}



