package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Page struct {
        Images []string
}

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Serve static files (HTML, CSS, JS)
    router.LoadHTMLGlob("static/*")
    //router.Static("/static", "./static")

    images := []string{"image1.jpg", "image2.jpg", "image3.jpg"}

    // Define routes
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Images": images,
        })
    })

    // Start server
    router.Run()
}