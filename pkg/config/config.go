package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	ListenAddress string `yaml:"listen_address"`
	HostKey       string `yaml:"host_key_path"`
	HostKeyType   string `yaml:"host_key_type"`
	Banner        string `yaml:"banner"`
	
	// Authentication settings
	AllowPasswordAuth bool `yaml:"allow_password_auth"`
	AllowKeyAuth      bool `yaml:"allow_key_auth"`
	
	// Forwarding settings
	ForwardEnabled  bool   `yaml:"forward_enabled"`
	ForwardHost     string `yaml:"forward_host"`
	ForwardPort     int    `yaml:"forward_port"`
	
	// Logging settings
	LogPath string `yaml:"log_path"`
}

// Load loads configuration from a file
func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{
		// Set defaults
		ListenAddress:    "0.0.0.0:2222",
		HostKeyType:      "rsa",
		AllowPasswordAuth: true,
		AllowKeyAuth:      false,
		ForwardEnabled:    false,
		LogPath:           "logs/honeygo.log",
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		ListenAddress:     "0.0.0.0:2222",
		HostKeyType:       "rsa",
		Banner:            "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.4",
		AllowPasswordAuth: true,
		AllowKeyAuth:      false,
		ForwardEnabled:    false,
		LogPath:           "logs/honeygo.log",
	}
}