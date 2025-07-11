package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Veera !!",
		})
	})
	r.Run("0.0.0.0:8080") // 👈 Bind to all interfaces
}
