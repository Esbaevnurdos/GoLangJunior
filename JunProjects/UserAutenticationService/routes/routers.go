package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    // Public routes
    router.POST("/signup", signup)
    router.POST("/login", login)

    // Authenticated routes
    // authenticated := router.Group("/")
    // authenticated.Use(middlewares.Authenticate)
    // Add more authenticated routes here
}
