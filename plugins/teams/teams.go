package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// TeamsNotifier struct for the Microsoft Teams channel
type TeamsNotifier struct {
	webhookURL string
}

// Name returns the name of the notifier
func (t *TeamsNotifier) Name() string {
	return "Teams"
}

// Type returns the type of the notifier
func (t *TeamsNotifier) Type() string {
	return "teams"
}

// Notify sends a message to the Microsoft Teams webhook
func (t *TeamsNotifier) Notify(message *config.Message) error {
	if t.webhookURL == "" {
		return errors.New("webhook URL is not set")
	}

	payload := map[string]string{
		"text": message.Text,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(t.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification, received status code: %d", resp.StatusCode)
	}

	fmt.Println("Notification sent to Teams successfully")
	return nil
}

// New creates a new TeamsNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	webhookURL, ok := config["webhook_url"].(string)
	if !ok || webhookURL == "" {
		return nil, errors.New("missing or invalid webhook URL")
	}
	return &TeamsNotifier{webhookURL: webhookURL}, nil
}
