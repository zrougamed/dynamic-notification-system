package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// DiscordNotifier struct for the Discord channel
type DiscordNotifier struct {
	webhookURL string
}

// Name returns the name of the notifier
func (d *DiscordNotifier) Name() string {
	return "Discord"
}

// Type returns the type of the notifier
func (d *DiscordNotifier) Type() string {
	return "discord"
}

// Notify sends a message to the Discord webhook
func (d *DiscordNotifier) Notify(message *config.Message) error {
	if d.webhookURL == "" {
		return errors.New("webhook URL is not set")
	}

	// Create payload
	payload := map[string]string{
		"content": message.Text,
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Send POST request to the Discord webhook
	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send notification, received status code: %d", resp.StatusCode)
	}

	fmt.Println("Notification sent to Discord successfully")
	return nil
}

// New creates a new DiscordNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	webhookURL, ok := config["webhook_url"].(string)
	if !ok || webhookURL == "" {
		return nil, errors.New("missing or invalid webhook URL")
	}
	return &DiscordNotifier{webhookURL: webhookURL}, nil
}
