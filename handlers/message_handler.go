package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/websocket-chat/models"
	"gorm.io/gorm"
)

type MessageHandler struct {
	DB *gorm.DB
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	recipientID, err := strconv.ParseUint(c.Param("recipient_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient ID"})
		return
	}

	currentUserID := c.GetUint("userID")

	var messages []models.Message

	h.DB.Where("(sender_id = ? AND recipient_id = ?) OR (sender_id = ? AND recipient_id = ?)",
		currentUserID, recipientID, recipientID, currentUserID).
		Order("created_at asc").
		Find(&messages)

	c.JSON(http.StatusOK, messages)
}
