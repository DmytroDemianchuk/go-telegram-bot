package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "6757927413:AAE0qUx2Ek3w8ch1jRPb6x7e7UlwFpQaRKA"

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Delete webhook before starting long polling
	err = bot.DeleteWebhook(nil)
	if err != nil {
		fmt.Println("Error deleting webhook:", err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error getting bot info:", err)
		os.Exit(1)
	}
	fmt.Printf("Bot User: %+v\n", botUser)

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Println("Error getting updates:", err)
		os.Exit(1)
	}
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage.
			// Send a message to sender with the same text (echo bot).
			// (https://core.telegram.org/bots/api#sendmessage)
			sentMessage, err := bot.SendMessage(
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)
			if err != nil {
				fmt.Println("Error sending message:", err)
				continue
			}

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
