package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Notifier interface {
	Name() string
	Notify(message string) error
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
}

type Config struct {
	Channels map[string]ChannelConfig `yaml:"channels"`
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
