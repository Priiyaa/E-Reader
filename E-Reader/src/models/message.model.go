package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID   string             `bson:"senderId" json:"senderId"`
	ReceiverID string             `bson:"receiverId" json:"receiverId"`
	Text       string             `bson:"text" json:"text"`
	Image      string             `bson:"image" json:"image"`
}
