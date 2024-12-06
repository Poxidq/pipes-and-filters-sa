package api

import (
	"fmt"
	"net/http"
	"paf/config"
	"paf/shared"

	"github.com/gin-gonic/gin"
)

var ApiService Api = Api{}

type Api struct {
	Next shared.Service
	Data chan shared.Message
}

func handlePostMessage(c *gin.Context) {
	var msg shared.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'user' and 'message' fields are required and must be valid JSON.",
		})
		return
	}
	ApiService.CallNext(msg)

	c.JSON(http.StatusOK, gin.H{
		"message": "Message successfully sent to next service.",
	})
}

func (api *Api) Execute(p *shared.Process) {
	router := gin.Default()
	router.POST("/messages", handlePostMessage)
	fmt.Printf("config.CFG.ApiPort: %s\n", config.CFG.ApiPort)
	runString := fmt.Sprintf(":%s", config.CFG.ApiPort)

	router.Run(runString)
}

func (c *Api) SetNext(next shared.Service) {
	c.Next = next
}

func (c *Api) CallNext(msg shared.Message) {
	if c.Next != nil {
		go c.Next.Execute(&shared.Process{Message: msg})
	}
}
