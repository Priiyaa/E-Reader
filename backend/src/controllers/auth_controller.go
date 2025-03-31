package controllers

import (
	"backend/src/lib"
	db "backend/src/lib"
	"backend/src/models"
	"backend/utils"
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func getUserCollection() *mongo.Collection {
	if db.DB == nil {
		log.Fatal("Database connection is not initialized!")
	}
	if userCollection == nil {
		userCollection = db.DB.Collection("users")
	}
	return userCollection
}

func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := getUserCollection()

	// Check if email exists
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// **Assign MongoDB ObjectID Correctly**
	newUser := models.User{
		ID:         primitive.NewObjectID(),
		FullName:   input.FullName,
		Email:      input.Email,
		Password:   string(hashedPassword),
		ProfilePic: "",
	}

	// Insert into MongoDB
	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT Token
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

	// **Find user by email**
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// **Verify password correctly**
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// **Generate JWT Token**
	token, _ := utils.GenerateToken(user.ID.Hex())

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

func Logout(c *gin.Context) {
	// Clear the JWT cookie by setting it to an empty value with an immediate expiry
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func UploadProfilePic(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
		return
	}

	// Save to temporary path
	tempFilePath := "temp-upload.jpg"
	out, err := os.Create(tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}
	defer out.Close()
	io.Copy(out, file)

	// Upload to Cloudinary
	imageURL, err := lib.UploadImage(tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}
