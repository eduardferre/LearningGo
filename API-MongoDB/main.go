package main

import (
	"log"

	// Gin API framework
	"github.com/gin-gonic/gin"
)

const serverURI = "localhost:8080"

var mongoClient *MongoClient

func init() {
	client, err := NewMongoStore()

	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	log.Print("Connected to MongoDB")
	mongoClient = client
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(serverURI)
}
