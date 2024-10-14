package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Telegram API URL with your Bot Token
var telegramBotToken string = os.Getenv("TELEGRAM_BOT_TOKEN")
var telegramAPIURL string = "https://api.telegram.org/bot" + telegramBotToken + "/sendMessage"

// Chat ID of the user or group where you want to send the message
var chatID string = os.Getenv("TELEGRAM_CHANNEL_ID")

// Struct to define the message payload for the Telegram API

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

var Telegram_msgs []string

// Function to send a message to the Telegram bot
func sendMessageToTelegram(message string) error {
	// Create the message payload
	Telegram_msgs = append(Telegram_msgs, message)
	fmt.Println("this is the message: ", message)
	if len(Telegram_msgs) == 10 {
		messages := strings.Join(Telegram_msgs, "\n")
		msg := TelegramMessage{
			ChatID: chatID,
			Text:   messages,
		}
		fmt.Println("sendMessageToTelegram received: ", msg)
		// Marshal the message payload to JSON
		payloadBytes, err := json.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed to marshal message: %w", err)
		}

		// Create a new HTTP POST request
		resp, err := http.Post(telegramAPIURL, "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}
		defer resp.Body.Close()

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			bodyString := string(bodyBytes)
			return fmt.Errorf("failed to send message, status code: %d, response: %s, response body: %s", resp.StatusCode, resp.Status, bodyString)
		}
		Telegram_msgs = []string{}
		fmt.Println("Message batch sent. Telegram_msgs array cleared.")
	}

	return nil
}
