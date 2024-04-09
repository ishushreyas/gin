package main

import (
            "context"
        "flag"
        "fmt"
        "html/template"
        "io"
        "log"
        "net/http"
        "os"
        "path/filepath"

       // "github.com/google/generative-ai-go/genai"
       // "google.golang.org/api/iterator"
      //  "google.golang.org/api/option"
    "github.com/gin-gonic/gin"
)

var apiKey = "AIzaSyBKLT-XRD7a0owyqqOIpsJjkU1SIlIHfO0"
/*
func generateHandler(w http.ResponseWriter, r *http.Request, model *genai.GenerativeModel) {
        if apiKey == "TODO" {
                http.Error(w, "Error: To get started, get an API key at https://makersuite.google.com/app/apikey and enter it in cmd/web/main.go and then hard restart the preview", http.StatusInternalServerError)
                return
        }

        image, prompt := r.FormValue("chosen-image"), r.FormValue("prompt")
        contents, err := os.ReadFile(filepath.Join("static", "images", filepath.Base(image)))
        if err != nil {
                log.Printf("Unable to read image %s: %v\n", image, err)
                http.Error(w, "Error: unable to generate content", http.StatusInternalServerError)
                return
        }

        // Generate the response and aggregate the streamed response.
        iter := model.GenerateContentStream(r.Context(), genai.Text(prompt), genai.ImageData("jpeg", contents))
        for {
                resp, err := iter.Next()
                if err == iterator.Done {
                        break
                }
                if err != nil {
                        log.Printf("Error generating content: %v\n", err)
                        http.Error(w, "Error: unable to generate content", http.StatusInternalServerError)
                        return
                }
                if resp == nil {
                        continue
                }
                for _, cand := range resp.Candidates {
                        if cand.Content != nil {
                                for _, part := range cand.Content.Parts {
                                        fmt.Fprint(w, part)
                                }
                        }
                }
        }
}*/

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Serve static files (HTML, CSS, JS)
    router.LoadHTMLGlob("templates/*")
    router.Static("/static", "./static")

    images := []string{"image1.jpg", "image2.jpg", "image3.jpg"}

    // Define routes
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Images": images,
        })
    })

    router.POST("/api/generate", func(w http.ResponseWriter, r *http.Request) { /*generateHandler(w, r, model) */})

    // Start server
    router.Run()
}
