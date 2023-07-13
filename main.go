package main

import (
	"auth-client/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user-token", server.GetUserToken)
	r.POST("/list", server.ListUser)
	r.Run(":3000")
}
