package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", controllers.Signup)
			auth.POST("/login", controllers.Login) // **New Login Route**
			auth.POST("/logout", controllers.Logout)
			auth.PUT("/update-profile", controllers.UploadProfilePic)
		}
	}
}
