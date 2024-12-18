package main

import (
	"dynamic-notification-system/config"
	"errors"
	"fmt"
	"net/smtp"
)

// SMTPNotifier struct for sending emails
type SMTPNotifier struct {
	host     string
	port     string
	username string
	password string
	to       string
}

// Name returns the name of the notifier
func (s *SMTPNotifier) Name() string {
	return "SMTP"
}

// Notify sends an email
func (s *SMTPNotifier) Notify(message *config.Message) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	to := []string{s.to}
	msg := []byte(fmt.Sprintf("Subject: Notification\n\n%s", message.Text))
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	err := smtp.SendMail(addr, auth, s.username, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Notification sent via SMTP successfully")
	return nil
}

// New creates a new SMTPNotifier instance
func New(config map[string]interface{}) (config.Notifier, error) {
	host, ok := config["host"].(string)
	port, ok2 := config["port"].(string)
	username, ok3 := config["username"].(string)
	password, ok4 := config["password"].(string)
	to, ok5 := config["to"].(string)

	if !(ok && ok2 && ok3 && ok4 && ok5) {
		return nil, errors.New("missing or invalid SMTP configuration")
	}

	return &SMTPNotifier{
		host:     host,
		port:     port,
		username: username,
		password: password,
		to:       to,
	}, nil
}
