package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RocketChatNotifier struct for Rocket.Chat
type RocketChatNotifier struct {
	webhookURL string
}

// Name returns the name of the notifier
func (r *RocketChatNotifier) Name() string {
	return "Rocket.Chat"
}

// Notify sends a message to a Rocket.Chat webhook
func (r *RocketChatNotifier) Notify(message string) error {
	if r.webhookURL == "" {
		return errors.New("webhook URL is not set")
	}

	payload := map[string]string{
		"text": message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(r.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification, received status code: %d", resp.StatusCode)
	}

	fmt.Println("Notification sent to Rocket.Chat successfully")
	return nil
}

// New creates a new RocketChatNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	webhookURL, ok := config["webhook_url"].(string)
	if !ok || webhookURL == "" {
		return nil, errors.New("missing or invalid webhook URL")
	}
	return &RocketChatNotifier{webhookURL: webhookURL}, nil
}
