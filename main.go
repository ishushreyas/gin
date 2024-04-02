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
    router.Static("/static", "./static")

            matches, err := filepath.Glob(filepath.Join("static", "images", "baked_goods_*.jpeg"))
        if err != nil {
                log.Printf("Error loading baked goods images: %v", err)
        }
        var page = &Page{Images: make([]string, len(matches))}
        for i, match := range matches {
                page.Images[i] = filepath.Base(match)
        }
        switch r.URL.Path {
        case "/":
                err = tmpl.Execute(w, page)
                if err != nil {
                        log.Printf("Template execution error: %v", err)
                }
        }

    // Define routes
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Images": "image1.jpg",
        })
    })

    // Start server
    router.Run()
}