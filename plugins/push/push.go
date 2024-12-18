package main

import (
	"dynamic-notification-system/config"
	"errors"
	"fmt"
)

// PushNotifier struct for push notifications
type PushNotifier struct {
	apiKey string
	device string
}

// Name returns the name of the notifier
func (p *PushNotifier) Name() string {
	return "Push Notification"
}

// Notify sends a push notification
func (p *PushNotifier) Notify(message *config.Message) error {
	fmt.Printf("Sending push notification to device %s with message: %s\n", p.device, message.Text)
	// WIP (e.g., Firebase, OneSignal)
	return nil
}

// New creates a new PushNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	apiKey, ok := config["api_key"].(string)
	device, ok2 := config["device"].(string)

	if !(ok && ok2) {
		return nil, errors.New("missing or invalid Push Notification configuration")
	}

	return &PushNotifier{
		apiKey: apiKey,
		device: device,
	}, nil
}
