// package controllers

// import (
// 	"backend/src/models"
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"golang.org/x/crypto/bcrypt"
// )

// // -----------------------------------------------------------------------------
// // Define the UserCollection interface for abstraction
// // -----------------------------------------------------------------------------
// type UserCollection interface {
// 	FindOne(ctx context.Context, filter interface{}) *mockResult
// 	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
// }

// // -----------------------------------------------------------------------------
// // MockUserCollection: Implements UserCollection for testing
// // -----------------------------------------------------------------------------
// type MockUserCollection struct {
// 	data map[string]models.User
// }

// func (m *MockUserCollection) FindOne(ctx context.Context, filter interface{}) *mockResult {
// 	email, ok := filter.(bson.M)["email"].(string)
// 	if !ok {
// 		return &mockResult{err: errors.New("invalid query format")}
// 	}

// 	user, exists := m.data[email]
// 	if !exists {
// 		log.Println("Mock DB: User not found ->", email) // Debugging log
// 		return &mockResult{err: errors.New("user not found")}
// 	}
// 	log.Println("Mock DB: User found ->", user.Email) // Debugging log
// 	return &mockResult{data: user}
// }

// func (m *MockUserCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
// 	return nil, nil
// }

// // -----------------------------------------------------------------------------
// // mockResult: Simulates MongoDB's SingleResult type
// // -----------------------------------------------------------------------------
// type mockResult struct {
// 	data models.User
// 	err  error
// }

// func (r *mockResult) Decode(v interface{}) error {
// 	if r.err != nil {
// 		return r.err
// 	}
// 	*v.(*models.User) = r.data
// 	return nil
// }

// // -----------------------------------------------------------------------------
// // Override getUserCollection for testing
// // -----------------------------------------------------------------------------
// var getUserCollection = func() UserCollection {
// 	// Generate bcrypt hash for "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

// 	mockUser := models.User{
// 		ID:         primitive.NewObjectID(),
// 		FullName:   "Test User",
// 		Email:      "test@example.com",
// 		Password:   string(hashedPassword), // Use the generated hash
// 		ProfilePic: "profile.jpg",
// 	}

// 	log.Println("Mock user created with email:", mockUser.Email) // Debugging log

// 	return &MockUserCollection{
// 		data: map[string]models.User{
// 			mockUser.Email: mockUser,
// 		},
// 	}
// }

// // -----------------------------------------------------------------------------
// // Login function (should be referenced in actual auth controller)
// // -----------------------------------------------------------------------------
// func Login(c *gin.Context) {
// 	var input struct {
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	collection := getUserCollection()

// 	// Find user by email
// 	var user models.User
// 	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
// 	if err != nil {
// 		log.Println("Login Error: User not found or wrong email") // Debugging log
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	// Verify password
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
// 		log.Println("Login Error: Password mismatch") // Debugging log
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	// Dummy token for test
// 	token := "dummy-token"

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": token,
// 		"user": gin.H{
// 			"id":         user.ID.Hex(),
// 			"fullName":   user.FullName,
// 			"email":      user.Email,
// 			"profilePic": user.ProfilePic,
// 		},
// 	})
// }

// // -----------------------------------------------------------------------------
// // TestLogin: Unit tests for the Login handler
// // -----------------------------------------------------------------------------
// func TestLogin(t *testing.T) {
// 	// Set Gin to test mode
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()

// 	// Register login route with the test Login function
// 	router.POST("/login", Login)

// 	t.Run("✅ Successful login", func(t *testing.T) {
// 		input := map[string]string{
// 			"email":    "test@example.com",
// 			"password": "password123",
// 		}
// 		jsonInput, _ := json.Marshal(input)
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp := httptest.NewRecorder()
// 		router.ServeHTTP(resp, req)

// 		assert.Equal(t, http.StatusOK, resp.Code)
// 	})

// 	t.Run("❌ Invalid email", func(t *testing.T) {
// 		input := map[string]string{
// 			"email":    "wrong@example.com",
// 			"password": "password123",
// 		}
// 		jsonInput, _ := json.Marshal(input)
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp := httptest.NewRecorder()
// 		router.ServeHTTP(resp, req)

// 		assert.Equal(t, http.StatusUnauthorized, resp.Code)
// 	})

// 	t.Run("❌ Invalid password", func(t *testing.T) {
// 		input := map[string]string{
// 			"email":    "test@example.com",
// 			"password": "wrongpassword",
// 		}
// 		jsonInput, _ := json.Marshal(input)
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp := httptest.NewRecorder()
// 		router.ServeHTTP(resp, req)

// 		assert.Equal(t, http.StatusUnauthorized, resp.Code)
// 	})

// 	t.Run("❌ Missing fields", func(t *testing.T) {
// 		input := map[string]string{
// 			"email": "test@example.com",
// 		}
// 		jsonInput, _ := json.Marshal(input)
// 		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
// 		req.Header.Set("Content-Type", "application/json")
// 		resp := httptest.NewRecorder()
// 		router.ServeHTTP(resp, req)

// 		assert.Equal(t, http.StatusBadRequest, resp.Code)
// 	})
// }

// login_test.go
package controllers

import (
	"backend/src/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login handler to be tested.
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	collection := getUserCollection()

	// Find user by email.
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify password.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Dummy token for test.
	token := "dummy-token"

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID.Hex(),
			"fullName":   user.FullName,
			"email":      user.Email,
			"profilePic": user.ProfilePic,
		},
	})
}

// TestLogin: Unit tests for the Login handler.
func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register login route.
	router.POST("/login", Login)

	t.Run("✅ Successful login", func(t *testing.T) {
		input := map[string]string{
			"email":    "test@example.com",
			"password": "password123",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("❌ Invalid email", func(t *testing.T) {
		input := map[string]string{
			"email":    "wrong@example.com",
			"password": "password123",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("❌ Invalid password", func(t *testing.T) {
		input := map[string]string{
			"email":    "test@example.com",
			"password": "wrongpassword",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	t.Run("❌ Missing fields", func(t *testing.T) {
		input := map[string]string{
			"email": "test@example.com",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
