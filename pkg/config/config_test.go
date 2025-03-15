package config

import (
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	
	if cfg.ListenAddress != "0.0.0.0:2222" {
		t.Errorf("Expected default ListenAddress to be 0.0.0.0:2222, got %s", cfg.ListenAddress)
	}
	
	if cfg.HostKeyType != "rsa" {
		t.Errorf("Expected default HostKeyType to be rsa, got %s", cfg.HostKeyType)
	}
	
	if !cfg.AllowPasswordAuth {
		t.Errorf("Expected default AllowPasswordAuth to be true")
	}
	
	if cfg.AllowKeyAuth {
		t.Errorf("Expected default AllowKeyAuth to be false")
	}
}

func TestLoad(t *testing.T) {
	// Create a temporary config file
	content := `listen_address: "127.0.0.1:2222"
host_key_path: "test_key"
host_key_type: "ed25519"
banner: "Test Banner"
allow_password_auth: false
allow_key_auth: true
forward_enabled: true
forward_host: "localhost"
forward_port: 8888
log_path: "test.log"`
	
	tmpfile, err := os.CreateTemp("", "config_test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	
	// Load the config
	cfg, err := Load(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	
	// Check values
	if cfg.ListenAddress != "127.0.0.1:2222" {
		t.Errorf("Expected ListenAddress to be 127.0.0.1:2222, got %s", cfg.ListenAddress)
	}
	
	if cfg.HostKey != "test_key" {
		t.Errorf("Expected HostKey to be test_key, got %s", cfg.HostKey)
	}
	
	if cfg.HostKeyType != "ed25519" {
		t.Errorf("Expected HostKeyType to be ed25519, got %s", cfg.HostKeyType)
	}
	
	if cfg.Banner != "Test Banner" {
		t.Errorf("Expected Banner to be Test Banner, got %s", cfg.Banner)
	}
	
	if cfg.AllowPasswordAuth {
		t.Errorf("Expected AllowPasswordAuth to be false")
	}
	
	if !cfg.AllowKeyAuth {
		t.Errorf("Expected AllowKeyAuth to be true")
	}
	
	if !cfg.ForwardEnabled {
		t.Errorf("Expected ForwardEnabled to be true")
	}
	
	if cfg.ForwardHost != "localhost" {
		t.Errorf("Expected ForwardHost to be localhost, got %s", cfg.ForwardHost)
	}
	
	if cfg.ForwardPort != 8888 {
		t.Errorf("Expected ForwardPort to be 8888, got %d", cfg.ForwardPort)
	}
	
	if cfg.LogPath != "test.log" {
		t.Errorf("Expected LogPath to be test.log, got %s", cfg.LogPath)
	}
}

func TestLoadNonExistentFile(t *testing.T) {
	_, err := Load("non_existent_file.yaml")
	if err == nil {
		t.Errorf("Expected error when loading non-existent file, got nil")
	}
}