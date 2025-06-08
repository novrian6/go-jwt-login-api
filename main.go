package main

import (
	"go-jwt-login-api/controllers"
	"go-jwt-login-api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vue dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public route
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/").Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
	}

	r.Run(":8080")
}
