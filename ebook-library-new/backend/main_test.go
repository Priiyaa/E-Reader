// main_test.go
package main

import (
	"bytes"
	"context"

	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock S3 Client for Upload
type MockUploader struct {
	UploadFunc func(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func (m *MockUploader) Upload(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return m.UploadFunc(ctx, input)
}

func TestHandleUpload(t *testing.T) {
	// Create a mock S3 uploader with a mocked upload function
	mockUploader := &MockUploader{
		UploadFunc: func(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			// Check that the file name starts with "testfile.pdf"
			assert.True(t, string(*input.Key)[:len("testfile.pdf")] == "testfile.pdf")
			return &s3.PutObjectOutput{}, nil
		},
	}

	// Create a temporary file to upload
	tmpFile, err := os.CreateTemp("", "testfile.pdf")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test
	tmpFile.Write([]byte("some dummy file content"))
	tmpFile.Close()

	// Create a buffer to store the multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the multipart form data with the field name "pdf"
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to open temporary file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("pdf", tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, err = io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file content: %v", err)
	}
	_, err = bytes.NewReader([]byte("some dummy file content")).WriteTo(part)
	if err != nil {
		t.Fatalf("Failed to write content to form file: %v", err)
	}

	// Close the writer to set the correct content type
	writer.Close()

	// Set up Gin context with the mock uploader
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/upload", func(c *gin.Context) {
		handleUpload(c, mockUploader) // Use the mock uploader here
	})

	// Create a request to upload the file
	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // Set correct content type

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status
	assert.Equal(t, 200, w.Code)

	// Check the response body contains the correct URL and file name pattern
	assert.Contains(t, w.Body.String(), "https://books-uploaded.s3.amazonaws.com/")
	assert.Contains(t, w.Body.String(), "testfile.pdf") // Check that the file name starts with "testfile.pdf"
}
