package routes

import "github.com/gin-gonic/gin"

func HealthRoutes(router *gin.Engine) {
	status := router.Group("/health")
	{
		status.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Server is UP!",
			})
		})

	}
}