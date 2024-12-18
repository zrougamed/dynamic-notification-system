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

// NfytNotifier handles sending notifications to nfyt
type NfytNotifier struct {
	apiKey string
	Topic  string
	Server string
}

// Name returns the name of the notifier
func (n *NfytNotifier) Name() string {
	return "Nfyt"
}

// Helper function to validate the Actions
func validateActions(actions []interface{}) error {
	validActions := map[string]bool{"view": true, "broadcast": true, "http": true}

	for _, action := range actions {
		actMap, ok := action.(map[string]string)
		if !ok {
			return errors.New("invalid action format")
		}
		if !validActions[actMap["action"]] {
			return fmt.Errorf("invalid action type: %s. Valid values are 'view', 'broadcast', 'http'", actMap["action"])
		}
	}
	return nil
}

// Notify sends a notification via nfyt using the Message object
func (n *NfytNotifier) Notify(message *config.Message) error {
	if n.apiKey == "" {
		return errors.New("missing API key for nfyt")
	}
	if n.Topic == "" {
		return errors.New("missing topic for nfyt")
	}

	// Validate priority
	if message.Priority < 1 || message.Priority > 5 {
		return fmt.Errorf("invalid priority value: %d. Must be between 1 and 5", message.Priority)
	}

	// Validate Actions
	err := validateActions(message.Actions)
	if err != nil {
		return fmt.Errorf("action validation failed: %v", err)
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
		return fmt.Errorf("nfyt API returned status: %s\nHeaders: %v\nBody: %s",
			resp.Status, resp.Header, string(body))
	}
	fmt.Printf("Message sent to nfyt topic '%s': %s\n", n.Topic, message.Text)
	return nil
}

// New is the constructor function required by the plugin system
func New(config map[string]interface{}) (config.Notifier, error) {
	apiKey, ok := config["api_key"].(string)
	if !ok || apiKey == "" {
		return nil, errors.New("invalid or missing API key for nfyt")
	}

	topic, ok := config["topic"].(string)
	if !ok || topic == "" {
		return nil, errors.New("invalid or missing topic for nfyt")
	}

	server, ok := config["server"].(string)
	if !ok || server == "" {
		return nil, errors.New("invalid or missing server for nfyt")
	}

	return &NfytNotifier{
		apiKey: apiKey,
		Topic:  topic,
		Server: server,
	}, nil
}
