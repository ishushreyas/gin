package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/templates/index.tmpl.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "index.tmpl.html", gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
