package main

import (
	"log"

	"github.com/eduardferre/learninggo/api-mongodb/db"
	"github.com/eduardferre/learninggo/api-mongodb/routes"
	"go.mongodb.org/mongo-driver/mongo"

	// Gin API framework
	"github.com/gin-gonic/gin"
)

const serverURI = "localhost:8080"

var mongoClient *mongo.Client
var mongoDb *mongo.Database

func init() {
	client, err := db.NewMongoStore()

	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	log.Print("Connected to MongoDB")
	mongoClient = client
	mongoDb = mongoClient.Database("test_mongo_go")
}

func main() {
	router := gin.Default()

	routes.AuthRoutes(router, mongoDb)
	routes.HealthRoutes(router)

	router.Run(serverURI)
}
