package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// SlackNotifier struct for the Slack channel
type SlackNotifier struct {
	webhookURL string
}

// Name returns the name of the notifier
func (s *SlackNotifier) Name() string {
	return "Slack"
}

// Notify sends a message to the Slack webhook
func (s *SlackNotifier) Notify(message *config.Message) error {
	if s.webhookURL == "" {
		return errors.New("webhook URL is not set")
	}

	payload := map[string]string{
		"text": message.Text,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification, received status code: %d", resp.StatusCode)
	}

	fmt.Println("Notification sent to Slack successfully")
	return nil
}

// New creates a new SlackNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	webhookURL, ok := config["webhook_url"].(string)
	if !ok || webhookURL == "" {
		return nil, errors.New("missing or invalid webhook URL")
	}
	return &SlackNotifier{webhookURL: webhookURL}, nil
}
