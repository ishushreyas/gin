package main

import (
	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "get method successful",
		})
	})
	r.Run()
}
