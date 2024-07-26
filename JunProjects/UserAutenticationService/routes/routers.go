package routes

import (
	middlewares "go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    // Public routes
    router.POST("/signup", signUp)
    router.POST("/login", login)

    // Authenticated routes
    authenticated := router.Group("/")
    authenticated.Use(middlewares.Authenticate)
    // Add more authenticated routes here
}
