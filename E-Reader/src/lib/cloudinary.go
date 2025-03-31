package lib

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

// Cloudinary instance
var Cloudinary *cloudinary.Cloudinary

// InitCloudinary initializes Cloudinary with credentials from .env
func InitCloudinary() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file")
	}

	// Get environment variables
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Cloudinary environment variables are missing!")
	}

	// Initialize Cloudinary
	Cloudinary, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}

	log.Println("✅ Cloudinary initialized successfully!")
}
