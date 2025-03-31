package models

// import "go.mongodb.org/mongo-driver/bson/primitive"
import (
	"context"

	db "backend/src/lib"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID   string             `bson:"senderId" json:"senderId"`
	ReceiverID string             `bson:"receiverId" json:"receiverId"`
	Text       string             `bson:"text" json:"text"`
	Image      string             `bson:"image" json:"image"`
}

func SaveMessageToDB(msg Message) error {
	collection := db.DB.Collection("messages")
	_, err := collection.InsertOne(context.TODO(), msg)
	return err
}
