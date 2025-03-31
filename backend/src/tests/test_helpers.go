// test_helpers.go
package controllers

import (
	"context"
	"errors"
	"log"

	"backend/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// UserCollection abstracts database operations.
type UserCollection interface {
	FindOne(ctx context.Context, filter interface{}) *mockResult
	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
}

// mockResult simulates MongoDB's SingleResult type.
type mockResult struct {
	data models.User
	err  error
}

func (r *mockResult) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	*v.(*models.User) = r.data
	return nil
}

// MockUserCollection is a common mock for testing.
type MockUserCollection struct {
	data map[string]models.User
}

func (m *MockUserCollection) FindOne(ctx context.Context, filter interface{}) *mockResult {
	email, ok := filter.(bson.M)["email"].(string)
	if !ok {
		return &mockResult{err: errors.New("invalid query format")}
	}

	user, exists := m.data[email]
	if !exists {
		log.Println("Mock DB: User not found ->", email)
		return &mockResult{err: errors.New("user not found")}
	}
	log.Println("Mock DB: User found ->", user.Email)
	return &mockResult{data: user}
}

func (m *MockUserCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	// For tests like Login, InsertOne is not used.
	return nil, nil
}

// getUserCollection is a function that returns a preconfigured test collection.
var getUserCollection = func() UserCollection {
	// Generate bcrypt hash for "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	mockUser := models.User{
		ID:         primitive.NewObjectID(),
		FullName:   "Test User",
		Email:      "test@example.com",
		Password:   string(hashedPassword),
		ProfilePic: "profile.jpg",
	}

	log.Println("Mock user created with email:", mockUser.Email)

	return &MockUserCollection{
		data: map[string]models.User{
			mockUser.Email: mockUser,
		},
	}
}
