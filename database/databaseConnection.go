package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// func DBinstance() *mongo.Client {
// 	err := godotenv.Load()
//   	if err != nil {
//     	log.Fatal("Error loading .env file")
//     }
// 	mongoDb := os.Getenv("MONGODB_URI")
// 	fmt.Println("mongo_url",mongoDb)

// 	client, err := mongo.Connect(options.Client().ApplyURI(mongoDb))

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

	
// 	fmt.Println("Connected to MongoDB")
// 	return client
// }


func DBinstance() *mongo.Client {

	// Load .env ONLY if present (local dev)
	_ = godotenv.Load()

	mongoDb := os.Getenv("MONGODB_URI")
	if mongoDb == "" {
		log.Fatal("MONGODB_URI not set")
	}

	fmt.Println("mongo_url:", mongoDb)

	client, err := mongo.Connect(options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}


var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("vender").Collection(collectionName)
}


// import (
// 	"context"
// 	"log"
// 	"os"
// 	"time"

// 	"go.mongodb.org/mongo-driver/v2/mongo"
// 	"go.mongodb.org/mongo-driver/v2/mongo/options"
// )

// var Client *mongo.Client

// func ConnectDB() {
// 	mongoURI := os.Getenv("MONGODB_URI")
// 	if mongoURI == "" {
// 		log.Fatal("MONGODB_URI is not set")
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Ping DB
// 	if err := client.Ping(ctx, nil); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("✅ Connected to MongoDB")
// 	Client = client
// }

// func OpenCollection(collectionName string) *mongo.Collection {
// 	return Client.Database("vendercollection").Collection(collectionName)
// }
