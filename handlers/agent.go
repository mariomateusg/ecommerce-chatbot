package handlers

import (
	"context"
	"ecommerce-bot/bot"
	"ecommerce-bot/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentHandler struct {
	ctx     context.Context
	chatbot *bot.AgentCase
}

func NewAgentHandler(ctx context.Context, chat *bot.AgentCase) *AgentHandler {
	return &AgentHandler{
		ctx:     ctx,
		chatbot: chat,
	}
}

// swagger:route GET /greetings
//
// Get greeting message
//
// Responses:
// - 200: string

func (handler *AgentHandler) GreetingsHandler(c *gin.Context) {

	greetings := handler.chatbot.Greetings()
	c.JSON(http.StatusOK, greetings)

}

// swagger:route POST /reply
//
// Get reply from bot
//
// Responses:
// - 200: ReplyOutput
func (handler *AgentHandler) ReplyHandler(c *gin.Context) {

	var replyEntry models.ReplyEntry
	if err := c.ShouldBindJSON(&replyEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	reply := handler.chatbot.Reply(replyEntry.Query)

	replyOut := models.ReplyOutput{
		Result: reply,
	}

	c.JSON(http.StatusOK, replyOut)
}
