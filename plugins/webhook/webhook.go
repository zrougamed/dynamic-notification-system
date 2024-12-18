package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// WebhookNotifier struct for generic webhook notifications
type WebhookNotifier struct {
	url string
}

// Name returns the name of the notifier
func (w *WebhookNotifier) Name() string {
	return "Webhook"
}

// Notify sends a message to a generic webhook
func (w *WebhookNotifier) Notify(message *config.Message) error {
	if w.url == "" {
		return errors.New("webhook URL is not set")
	}

	payload := map[string]string{
		"message": message.Text,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(w.url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send notification, received status code: %d", resp.StatusCode)
	}

	fmt.Println("Notification sent to Webhook successfully")
	return nil
}

// New creates a new WebhookNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	url, ok := config["url"].(string)
	if !ok || url == "" {
		return nil, errors.New("missing or invalid webhook URL")
	}
	return &WebhookNotifier{url: url}, nil
}
