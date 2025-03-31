package main

import (
	"backend/src/lib"
	"backend/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB and Cloudinary
	lib.ConnectDB()
	lib.InitCloudinary()

	r := gin.Default()

	// ✅ Register Routes
	routes.RegisterAuthRoutes(r)
	routes.RegisterMessageRoutes(r) // ✅ Ensures `/api/messages/users` is available

	r.Run(":5002") // ✅ Start server on port 5002
}
