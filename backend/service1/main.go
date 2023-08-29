package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin with the default middleware
    router := gin.Default()

    // Define a route
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, Gin!",
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}
