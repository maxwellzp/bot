package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/maxwellzp/bot/internal/service/product"
)

// var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {

		if panicValue := recover(); panicValue != nil {
			fmt.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data: "+update.CallbackQuery.Data,
		)
		c.bot.Send(msg)
		return
	}

	if update.Message != nil { // If we got a message

		// command, ok := registeredCommands[update.Message.Command()]
		// if ok {
		// 	command(c, update.Message)
		// } else {
		// 	c.Default(update.Message)
		// }

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}
