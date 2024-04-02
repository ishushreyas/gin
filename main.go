package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Serve static files (HTML, CSS, JS)
    router.Static("/static", "./static")

    // Define routes
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    // Start server
    router.Run()
}