package websocket

import (
	db "backend/src/lib"
	"backend/src/models" // For both user and message models
	"context"
	"encoding/json"
	"fmt"
	"sync"

	// "your_project/models" // Replace with actual path
	// "your_project/database" // Replace with actual path

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var userSocketMap = make(map[string]*websocket.Conn)
var mutex = &sync.Mutex{}

func broadcastOnlineUsers() {
	mutex.Lock()
	defer mutex.Unlock()

	var onlineUsers []string
	for userId := range userSocketMap {
		onlineUsers = append(onlineUsers, userId)
	}

	for _, conn := range userSocketMap {
		err := conn.WriteJSON(map[string]interface{}{
			"type":  "getOnlineUsers",
			"users": onlineUsers,
		})
		if err != nil {
			fmt.Println("Error broadcasting:", err)
		}
	}
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	userId := c.Query("userId")
	if userId == "" {
		conn.Close()
		return
	}

	// Register user
	mutex.Lock()
	userSocketMap[userId] = conn
	mutex.Unlock()

	fmt.Println("User connected:", userId)
	broadcastOnlineUsers()

	defer func() {
		conn.Close()
		mutex.Lock()
		delete(userSocketMap, userId)
		mutex.Unlock()
		fmt.Println("User disconnected:", userId)
		broadcastOnlineUsers()
	}()

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		msgType, ok := msg["type"].(string)
		if !ok {
			fmt.Println("Invalid message format")
			continue
		}

		switch msgType {
		case "chatMessage":
			data, _ := json.Marshal(msg["data"])
			var message models.Message
			if err := json.Unmarshal(data, &message); err != nil {
				fmt.Println("Invalid message format:", err)
				continue
			}

			// Store in MongoDB
			if err := SaveMessageToDB(message); err != nil {
				fmt.Println("Error saving message:", err)
			}

			// Send to receiver if online
			if message.SenderID != userId {
				fmt.Println("Blocked message: Sender does not match socket user")
				continue
			}

			mutex.Lock()
			if receiverConn, ok := userSocketMap[message.ReceiverID]; ok {
				err := receiverConn.WriteJSON(map[string]interface{}{
					"type": "chatMessage",
					"data": message,
				})
				if err != nil {
					fmt.Println("Error sending message to receiver:", err)
				}
			}
			mutex.Unlock()
		default:
			fmt.Println("Unhandled message type:", msgType)
		}
	}
}

// Utility function to save a message in MongoDB
func SaveMessageToDB(message models.Message) error {
	collection := db.DB.Collection("messages")
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}
