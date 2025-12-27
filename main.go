package main

import (
//  "os"
//  "log"
 "github.com/gin-gonic/gin"
 //"go.mongodb.org/mongo-driver/v2/mongo"
 //"github.com/pharma999/vender/database"
 "github.com/pharma999/vender/routes"
 //"github.com/pharma999/vender/controller"
 //"github.com/joho/godotenv"
 "github.com/pharma999/vender/config"
)

//var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "vender")

func main() {
	//database.ConnectDB()
	config.LoadConfig()
	// err := godotenv.Load()
  	// if err != nil {
    // 	log.Fatal("Error loading .env file")
    // }
	port := config.GetEnv("PORT")

	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.VenderRoutes(router)
	routes.VenderDataRoutes(router)
	routes.ProductDetailRouter(router)

	router.Run(":" + port)
}



