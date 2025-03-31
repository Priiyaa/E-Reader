package controllers

import (
	"backend/src/lib"
	"backend/src/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUsersForSidebar - Fetches all users except the logged-in user
func GetUsersForSidebar(c *gin.Context) {
	userID, exists := c.Get("userID") // Extract userID from middleware
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var users []models.User
	collection := lib.DB.Collection("users")

	// Query: Get all users except the logged-in user
	cursor, err := collection.Find(context.TODO(), bson.M{"_id": bson.M{"$ne": userID}}, options.Find().SetProjection(bson.M{"password": 0}))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetMessages - Fetch chat messages between two users
func GetMessages(c *gin.Context) {
	authUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	partnerID := c.Param("id")

	if partnerID == authUserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot message yourself"})
		return
	}

	collection := lib.DB.Collection("messages")

	filter := bson.M{
		"$or": []bson.M{
			{"senderId": authUserID, "receiverId": partnerID},
			{"senderId": partnerID, "receiverId": authUserID},
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(context.TODO())

	var messages []models.Message
	if err := cursor.All(context.TODO(), &messages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// SendMessage - Handles sending a message
func SendMessage(c *gin.Context) {
	authUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	receiverID := c.Param("id")
	if receiverID == authUserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot send a message to yourself"})
		return
	}

	var input struct {
		Text  string `json:"text"`
		Image string `json:"image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newMessage := models.Message{
		ID:         primitive.NewObjectID(),
		SenderID:   authUserID.(string),
		ReceiverID: receiverID,
		Text:       input.Text,
		Image:      input.Image,
	}

	collection := lib.DB.Collection("messages")
	_, err := collection.InsertOne(context.TODO(), newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, newMessage)
}
