package main

import (
	"context"
	"ecommerce-bot/bot"
	"ecommerce-bot/handlers"
	"ecommerce-bot/internal/data"
	"log"

	"github.com/gin-gonic/gin"
)

var agentHandler *handlers.AgentHandler

func init() {
	ctx := context.Background()
	chatbot := bot.NewAgentCase()
	agentHandler = handlers.NewAgentHandler(ctx, chatbot)
}

func main() {

	router := gin.Default()

	router.GET("/greetings", agentHandler.GreetingsHandler)

	router.POST("/reply", agentHandler.ReplyHandler)

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	//chatbot := bot.NewAgentCase()

	//fmt.Println("Hola Mundo")

	//fmt.Println(chatbot.Greetings())

	//fmt.Println(chatbot.Reply("What is isomorphic go?"))

	router.Run()
}
