package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Serve static files (e.g., HTML, CSS, JS)
    r.Static("/static", "./bat")

    // Define routes
    r.GET("/", func(c *gin.Context) {
        c.File("./bat/main.html") // Serve your main HTML file
    })

    // Start the server
    r.Run(":8080") // Run on localhost:8080
}