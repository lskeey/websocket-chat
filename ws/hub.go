package ws

import (
	"encoding/json"
	"log"

	"github.com/lskeey/websocket-chat/models"
	"gorm.io/gorm"
)

type MessageWithSender struct {
	SenderID uint
	Message  []byte
}

type Hub struct {
	Clients    map[uint]*Client
	Broadcast  chan *MessageWithSender
	Register   chan *Client
	Unregister chan *Client
	DB         *gorm.DB
}

func NewHub(db *gorm.DB) *Hub {
	return &Hub{
		Clients:    make(map[uint]*Client),
		Broadcast:  make(chan *MessageWithSender),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		DB:         db,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserID] = client
			log.Printf("Client connected: UserID %d", client.UserID)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
				log.Printf("Client disconnected: UserID %d", client.UserID)
			}
		case messageData := <-h.Broadcast:
			var msg WsMessage
			if err := json.Unmarshal(messageData.Message, &msg); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			dbMessage := models.Message{
				SenderID:    messageData.SenderID,
				RecipientID: msg.RecipientID,
				Content:     msg.Content,
			}
			if err := h.DB.Create(&dbMessage).Error; err != nil {
				log.Printf("Failed to save message: %v", err)
				continue
			}

			recipientClient, ok := h.Clients[msg.RecipientID]
			if !ok {
				log.Printf("Recipient UserID %d is not online.", msg.RecipientID)
				continue
			}

			finalMessage, err := json.Marshal(dbMessage)
			if err != nil {
				log.Printf("Error marshaling message for delivery: %v", err)
				continue
			}

			select {
			case recipientClient.Send <- finalMessage:
			default:
				close(recipientClient.Send)
				delete(h.Clients, recipientClient.UserID)
			}
		}
	}
}
