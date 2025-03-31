package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	s3Client   *s3.Client
	uploader   *manager.Uploader
	bucketName = "books-uploaded"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize AWS S3 Client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Error initializing AWS config: %v", err)
	}

	// Assign the initialized client
	s3Client = s3.NewFromConfig(cfg)
	uploader = manager.NewUploader(s3Client) // Initialize uploader

	// Setup Gin app
	r := gin.Default()
	r.Use(cors.Default())

	// Routes for API endpoints
	r.POST("/upload", handleUpload)
	r.GET("/library", showLibrary)

	// Start server
	r.Run(":5000") // listen and serve on port 5000
}

// Handle file upload
func handleUpload(c *gin.Context) {
	// Get the uploaded file from the form
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	// Open the file
	f, openErr := file.Open()
	if openErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	// Upload file to S3
	_, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(file.Filename),
		Body:        f,
		ACL:         "public-read",
		ContentType: aws.String("application/pdf"),
	})
	if uploadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	fileURL := "https://" + bucketName + ".s3.amazonaws.com/" + file.Filename
	c.JSON(http.StatusOK, gin.H{"pdf_url": fileURL, "pdf_name": file.Filename})
}

// Show library with all uploaded PDFs
func showLibrary(c *gin.Context) {
	resp, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load library"})
		return
	}

	var fileLinks []string
	for _, item := range resp.Contents {
		fileLinks = append(fileLinks, "https://"+bucketName+".s3.amazonaws.com/"+*item.Key)
	}

	c.JSON(http.StatusOK, gin.H{"files": fileLinks})
}
