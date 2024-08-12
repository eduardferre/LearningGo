package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCollection *mongo.Collection

func AuthRoutes(router *gin.Engine, db *mongo.Database) {
	authCollection = db.Collection("auth")

	auth := router.Group("/auth")
	{
		auth.GET("/", getAllAuth)
	}
}

func getAllAuth(c *gin.Context) {
    cursor, err := authCollection.Find(context.TODO(), bson.D{{}})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var auth []bson.M
    if err = cursor.All(context.TODO(), &auth); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, auth)
}
