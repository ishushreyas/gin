package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Load HTML templates
    router.LoadHTMLGlob("templates/*")

    // Define route handlers
    router.GET("/", func(c *gin.Context) {
        // Render HTML template
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Shreyas",
        })
    })

    // Start server
    router.Run()
}
