package apimongodb

import (
	"log"

	// Gin API framework
	"github.com/gin-gonic/gin"
)

var mongoClient *MongoClient

func init() {
	client, err := NewMongoStore()

	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	log.Printf("Connected to MongoDB: %v", client)
	mongoClient = client
}

func main() {
	router := gin.Default()

	router.Run("localhost:8080")
}
