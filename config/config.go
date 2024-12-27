package config

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Message represents the structure of the notification message
type Message struct {
	Topic    string        `json:"topic,omitempty"`
	Title    string        `json:"title,omitempty"`
	Text     string        `json:"message,omitempty"`
	Tags     []string      `json:"tags,omitempty"`     // JSON array of tags
	Priority int           `json:"priority,omitempty"` // Integer 1=min, 3=default, 5=max
	Attach   string        `json:"attach,omitempty"`   // URL to a file
	Email    string        `json:"email,omitempty"`    // Email address for receiving notifications
	Actions  []interface{} `json:"actions,omitempty"`  // JSON Array of action buttons
}

// Implement sql.Scanner for Message
func (m *Message) Scan(value interface{}) error {
	// Ensure the value is a byte slice
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Message: expected []byte, got %T", value)
	}

	// Unmarshal JSON into the Message struct
	return json.Unmarshal(bytes, m)
}

// Implement driver.Valuer for inserting Message as JSON
func (m Message) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// ScheduledJob struct
type ScheduledJob struct {
	ID                 int          `json:"id,omitempty"` // omitempty for POST requests
	Name               string       `json:"name"`
	NotificationType   string       `json:"notification_type"`
	Recipient          string       `json:"recipient"`
	Message            Message      `json:"message"`
	ScheduleExpression string       `json:"schedule_expression"`
	LastRun            sql.NullTime `json:"last_run,omitempty"`
}

// InstantJob struct
type InstantJob struct {
	NotificationType string  `json:"notification_type"`
	Recipient        string  `json:"recipient"`
	Message          Message `json:"message"`
}

type Notifier interface {
	Name() string
	Type() string
	Notify(message *Message) error
}

type ChannelConfig struct {
	Enabled     bool   `yaml:"enabled"`
	WebhookURL  string `yaml:"webhook_url,omitempty"`
	URL         string `yaml:"url,omitempty"`
	ApiKey      string `yaml:"api_key,omitempty"`
	Host        string `yaml:"host,omitempty"`
	Port        string `yaml:"port,omitempty"`
	Username    string `yaml:"username,omitempty"`
	Password    string `yaml:"password,omitempty"`
	To          string `yaml:"to,omitempty"`
	Device      string `yaml:"device,omitempty"`
	ProviderAPI string `yaml:"provider_api,omitempty"`
	PhoneNumber string `yaml:"phone_number,omitempty"`
	ApiURL      string `yaml:"api_url,omitempty"`
	Topic       string `yaml:"topic,omitempty"`
	Server      string `yaml:"server,omitempty"`
}

type Config struct {
	Database  DatabaseConfig           `yaml:"database"`
	Channels  map[string]ChannelConfig `yaml:"channels"`
	Scheduler bool                     `yaml:"scheduler"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
