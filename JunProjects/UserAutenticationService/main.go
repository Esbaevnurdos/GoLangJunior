package main

import (
    "go-auth/database"
    "go-auth/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    database.InitDB()

    // Create a new Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Start the server
    router.Run(":8080")
}
