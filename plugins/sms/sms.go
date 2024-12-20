package main

import (
	"dynamic-notification-system/config"
	"errors"
	"fmt"
)

// SMSNotifier struct for SMS notifications
type SMSNotifier struct {
	providerAPI string
	apiKey      string
	phoneNumber string
}

// Name returns the name of the notifier
func (s *SMSNotifier) Name() string {
	return "SMS"
}

// Type returns the type of the notifier
func (s *SMSNotifier) Type() string {
	return "sms"
}

// Notify sends an SMS
func (s *SMSNotifier) Notify(message *config.Message) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.phoneNumber, message.Text)
	// WIP
	return nil
}

// New creates a new SMSNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	providerAPI, ok := config["provider_api"].(string)
	apiKey, ok2 := config["api_key"].(string)
	phoneNumber, ok3 := config["phone_number"].(string)

	if !(ok && ok2 && ok3) {
		return nil, errors.New("missing or invalid SMS configuration")
	}

	return &SMSNotifier{
		providerAPI: providerAPI,
		apiKey:      apiKey,
		phoneNumber: phoneNumber,
	}, nil
}
