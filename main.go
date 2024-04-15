package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/templates/index.tmpl.html")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl.html", gin.H{
			"image": string[]{"image1.jpg", "image2.jpg", "image3.jpg"},
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
