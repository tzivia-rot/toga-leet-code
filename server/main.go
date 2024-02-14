package main

import (
	"context"
	exerciseRouter "go-lenguage/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var mongoClient *mongo.Client
var collection *mongo.Collection

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	uri := os.Getenv("URI")
	if uri == "" {
		log.Fatal("URI not found in .env file")
	}

	if err := connectToMongoDB(uri); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
	collection = mongoClient.Database("togaLeetCode").Collection("exercises")
}

func main() {
	router := gin.Default()

	exerciseRouter.SetupRouter(router, collection)

	log.Fatal(router.Run(":8080"))
}

func connectToMongoDB(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}
