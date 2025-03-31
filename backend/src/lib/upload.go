package lib

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// UploadImage uploads an image to Cloudinary and returns the secure URL
func UploadImage(imagePath string) (string, error) {
	if Cloudinary == nil {
		log.Fatal("Cloudinary is not initialized")
	}

	// Provide valid UploadParams instead of `nil`
	uploadParams := uploader.UploadParams{
		Folder: "uploads", // Optional: Change folder name
	}

	// Upload image to Cloudinary
	uploadResult, err := Cloudinary.Upload.Upload(context.Background(), imagePath, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
