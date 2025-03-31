package routes

import (
	"backend/src/controllers"
	"backend/src/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterMessageRoutes registers message-related routes
func RegisterMessageRoutes(r *gin.Engine) {
	api := r.Group("/api/messages")      // ✅ This ensures the correct base URL
	api.Use(middleware.AuthMiddleware()) // ✅ Protects with JWT authentication

	api.GET("/users", controllers.GetUsersForSidebar) // ✅ Fetch users for sidebar
	api.GET("/:id", controllers.GetMessages)          // ✅ Get messages with a user
	api.POST("/send/:id", controllers.SendMessage)    // ✅ Send a message
}
