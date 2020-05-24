package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	hostname, _ := os.Hostname()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong v0.0.2 " + hostname,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
