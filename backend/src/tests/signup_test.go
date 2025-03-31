package controllers

import (
	"backend/src/models"
	"backend/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// MockUserCollectionSignup simulates the user collection for signup tests.
type MockUserCollectionSignup struct {
	data             map[string]models.User
	InsertOneFailure bool
}

func (m *MockUserCollectionSignup) FindOne(ctx context.Context, filter interface{}) *mockResult {
	email, ok := filter.(bson.M)["email"].(string)
	if !ok {
		return &mockResult{err: errors.New("invalid query format")}
	}

	user, exists := m.data[email]
	if exists {
		return &mockResult{data: user}
	}
	return &mockResult{err: errors.New("user not found")}
}

func (m *MockUserCollectionSignup) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	if m.InsertOneFailure {
		return nil, errors.New("failed to insert user")
	}

	user := document.(models.User)
	m.data[user.Email] = user
	return user.ID, nil
}

// Signup function to be tested.
func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := getUserCollection()

	// Check if email exists.
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Hash Password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	newUser := models.User{
		ID:         primitive.NewObjectID(),
		FullName:   input.FullName,
		Email:      input.Email,
		Password:   string(hashedPassword),
		ProfilePic: "",
	}

	// Insert into database.
	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token.
	token, _ := utils.GenerateToken(newUser.ID.Hex())

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
		"user": gin.H{
			"id":         newUser.ID.Hex(),
			"fullName":   newUser.FullName,
			"email":      newUser.Email,
			"profilePic": newUser.ProfilePic,
		},
	})
}

// TestSignup contains unit tests for the Signup handler.
func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", Signup)

	t.Run("✅ Successful signup", func(t *testing.T) {
		mockCollection := &MockUserCollectionSignup{
			data: make(map[string]models.User),
		}
		// Override getUserCollection for this test.
		getUserCollection = func() UserCollection { return mockCollection }

		input := models.User{
			FullName: "Test User",
			Email:    "newuser@example.com",
			Password: "password123",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("❌ Email already exists", func(t *testing.T) {
		// Prepare a collection with an existing user.
		existingUser := models.User{
			ID:       primitive.NewObjectID(),
			FullName: "Existing User",
			Email:    "existing@example.com",
			Password: "dummy", // password value is not checked in this scenario
		}
		mockCollection := &MockUserCollectionSignup{
			data: map[string]models.User{
				"existing@example.com": existingUser,
			},
		}
		getUserCollection = func() UserCollection { return mockCollection }

		input := models.User{
			FullName: "New User",
			Email:    "existing@example.com",
			Password: "password123",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("❌ Invalid request payload", func(t *testing.T) {
		// Supply malformed JSON (e.g., missing a closing quote or brace)
		invalidJSON := []byte(`{"FullName": "Test User", "email": "bademail", "password": }`)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("❌ Failed to insert user", func(t *testing.T) {
		mockCollection := &MockUserCollectionSignup{
			data:             make(map[string]models.User),
			InsertOneFailure: true,
		}
		getUserCollection = func() UserCollection { return mockCollection }

		input := models.User{
			FullName: "Test User",
			Email:    "failinsert@example.com",
			Password: "password123",
		}
		jsonInput, _ := json.Marshal(input)
		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonInput))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
