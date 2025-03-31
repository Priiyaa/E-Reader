package main

// import (
// 	"backend/src/lib"
// 	"backend/src/routes"
// 	"time"

// 	"github.com/gin-gonic/gin"

// 	"github.com/gin-contrib/cors"
// )

// func main() {
// 	// Initialize MongoDB and Cloudinary
// 	lib.ConnectDB()
// 	lib.InitCloudinary()

// 	r := gin.Default()

// 	r.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend origin
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	// ✅ Register Routes
// 	routes.RegisterAuthRoutes(r)
// 	routes.RegisterMessageRoutes(r) // ✅ Ensures `/api/messages/users` is available

// 	r.Run(":5002") // ✅ Start server on port 5002
// }

import (
	"backend/src/lib"
	"backend/src/routes"

	"backend/src/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	lib.ConnectDB()
	lib.InitCloudinary()

	r := gin.Default()
	// r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAuthRoutes(r)
	routes.RegisterMessageRoutes(r)

	// ✅ WebSocket route
	r.GET("/ws", websocket.HandleWebSocket)

	r.Run(":5002")
}
