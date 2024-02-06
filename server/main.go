
package main

import (
	"context"
	exerciseRouter "go-lenguage/routes"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var mongoClient *mongo.Client
var collection *mongo.Collection

func init() {
	if err := connectToMongoDB(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
	collection = mongoClient.Database("togaLeetCode").Collection("exercises")
}

func main() {
	router := gin.Default()

	exerciseRouter.SetupRouter(router, collection)

	log.Fatal(router.Run(":8080"))
}

func connectToMongoDB() error {
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
