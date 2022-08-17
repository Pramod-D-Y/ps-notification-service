package server

import (
	"github.com/gin-gonic/gin"
	"github.com/roppenlabs/ps-notification-service/internal/controller"
)

func StartServer() {
	router := gin.Default()
	router.GET("api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/notify", controller.CreateNotification)
	router.Run(":8080")
}
