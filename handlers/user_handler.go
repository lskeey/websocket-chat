package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/websocket-chat/models"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	usernameQuery := c.Query("username")
	if usernameQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username query parameter is required"})
		return
	}

	currentUserID := c.GetUint("userID")

	var users []models.User

	h.DB.Where("username LIKE ? AND id != ?", "%"+usernameQuery+"%", currentUserID).Find(&users)

	var response []UserResponse

	for _, user := range users {
		response = append(response, UserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	c.JSON(http.StatusOK, response)
}
