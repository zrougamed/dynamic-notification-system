package main

import (
	"dynamic-notification-system/config"
	"errors"
	"fmt"
)

// SignalNotifier struct for Signal messaging
type SignalNotifier struct {
	phoneNumber string
	apiURL      string
}

// Name returns the name of the notifier
func (s *SignalNotifier) Name() string {
	return "Signal"
}

// Notify sends a message via Signal
func (s *SignalNotifier) Notify(message string) error {
	fmt.Printf("Sending Signal message to %s: %s\n", s.phoneNumber, message)
	// WIP
	return nil
}

// New creates a new SignalNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	phoneNumber, ok := config["phone_number"].(string)
	apiURL, ok2 := config["api_url"].(string)

	if !(ok && ok2) {
		return nil, errors.New("missing or invalid Signal configuration")
	}

	return &SignalNotifier{
		phoneNumber: phoneNumber,
		apiURL:      apiURL,
	}, nil
}
