package main

import (
	"dynamic-notification-system/config"
	"errors"
	"fmt"
)

type TelegramNotifier struct {
	apiKey string
}

func (t *TelegramNotifier) Name() string {
	return "Telegram"
}

func (t *TelegramNotifier) Notify(message *config.Message) error {
	// WIP
	fmt.Printf("Sending message to Telegram: %s\n", message.Text)
	return nil
}

func New(config map[string]interface{}) (config.Notifier, error) {
	apiKey, ok := config["api_key"].(string)
	if !ok || apiKey == "" {
		return nil, errors.New("missing or invalid API key")
	}
	return &TelegramNotifier{apiKey: apiKey}, nil
}
