package main

import (
	"bytes"
	"dynamic-notification-system/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// NtfyNotifier handles sending notifications to ntfy
type NtfyNotifier struct {
	apiKey string
	Topic  string
	Server string
}

// Name returns the name of the notifier
func (n *NtfyNotifier) Name() string {
	return "Ntfy"
}

// Type returns the name of the notifier
func (n *NtfyNotifier) Type() string {
	return "ntfy"
}

// Notify sends a notification via ntfy using the Message object
func (n *NtfyNotifier) Notify(message *config.Message) error {
	if n.apiKey == "" {
		return errors.New("missing API key for ntfy")
	}
	if n.Topic == "" {
		return errors.New("missing topic for ntfy")
	}

	// Validate priority
	if message.Priority < 1 || message.Priority > 5 {
		return fmt.Errorf("invalid priority value: %d. Must be between 1 and 5", message.Priority)
	}

	// Adding the Topic to the Payload
	message.Topic = n.Topic

	// Marshal the message into JSON for POST body
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Convert JSON payload to string
	payloadString := string(payload)
	log.Printf("Payload: %s", payloadString)
	// Build the request
	req, err := http.NewRequest("POST", n.Server, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+n.apiKey)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	defer resp.Body.Close()

	// Debug response details in case of non-200 status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Read the response body
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("failed to read response body: %v", readErr)
		}

		// Print the full response for debugging
		return fmt.Errorf("ntfy API returned status: %s\nHeaders: %v\nBody: %s",
			resp.Status, resp.Header, string(body))
	}
	fmt.Printf("Message sent to ntfy topic '%s': %s\n", n.Topic, message.Title)
	return nil
}

// New is the constructor function required by the plugin system
func New(config map[string]interface{}) (config.Notifier, error) {
	apiKey, ok := config["api_key"].(string)
	if !ok || apiKey == "" {
		return nil, errors.New("invalid or missing API key for ntfy")
	}

	topic, ok := config["topic"].(string)
	if !ok || topic == "" {
		return nil, errors.New("invalid or missing topic for ntfy")
	}

	server, ok := config["server"].(string)
	if !ok || server == "" {
		return nil, errors.New("invalid or missing server for ntfy")
	}

	return &NtfyNotifier{
		apiKey: apiKey,
		Topic:  topic,
		Server: server,
	}, nil
}
