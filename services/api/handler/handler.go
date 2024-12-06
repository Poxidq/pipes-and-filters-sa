package handler

import (
	"net/http"
	"paf/shared"

	"github.com/gin-gonic/gin"
)

// HandlePostMessage handles the POST request to send messages
func HandlePostMessage(c *gin.Context) {
	var msg shared.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'user' and 'message' fields are required and must be valid JSON.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message successfully sent to next service.",
	})
}
