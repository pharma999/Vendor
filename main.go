// package main

// import (
//  "os"
//  "log"
//  "github.com/gin-gonic/gin"
//  //"go.mongodb.org/mongo-driver/v2/mongo"
//  //"github.com/pharma999/vender/database"
//  "github.com/pharma999/vender/routes"
//  //"github.com/pharma999/vender/controller"
//  "github.com/joho/godotenv"

// )

// //var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "vender")

// func main() {
// 	//database.ConnectDB()
// 	err := godotenv.Load()
//   	if err != nil {
//     	log.Fatal("Error loading .env file")
//     }
// 	port := os.Getenv("PORT")

// 	if port == ""{
// 		port = "8000"
// 	}

	
    

  
    

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.VenderRoutes(router)
// 	routes.VenderDataRoutes(router)
// 	routes.ProductDetailRouter(router)

// 	router.Run(":" + port)
// }



package main

import (
    "os"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/pharma999/vender/routes"
    "github.com/joho/godotenv"
)

func main() {

    // Load .env only if present (safe for all clouds)
    _ = godotenv.Load()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())

    routes.VenderRoutes(router)
    routes.VenderDataRoutes(router)
    routes.ProductDetailRouter(router)

    log.Println("Server running on port", port)
    router.Run(":" + port)
}
