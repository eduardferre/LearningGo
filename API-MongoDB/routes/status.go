package routes

import "github.com/gin-gonic/gin"

func StatusRoutes(router *gin.Engine) {
	status := router.Group("/status")
	{
		status.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Status OK!",
			})
		})

	}
}