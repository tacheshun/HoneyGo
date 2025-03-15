package ssh

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tacheshun/honeygo/pkg/config"
)

func TestNewServer(t *testing.T) {
	// Create a temporary directory for logs
	tempDir, err := os.MkdirTemp("", "server_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Create a test config
	cfg := &config.Config{
		ListenAddress:     "127.0.0.1:0", // Use port 0 to let the OS assign a port
		HostKeyType:       "rsa",
		Banner:            "SSH-2.0-Test",
		AllowPasswordAuth: true,
		AllowKeyAuth:      false,
		LogPath:           filepath.Join(tempDir, "test.log"),
	}
	
	// Create the server
	server, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Shutdown()
	
	// Verify that the server is not nil
	if server == nil {
		t.Fatalf("Server is nil")
	}
	
	// Verify that the server has the correct configuration
	if server.config != cfg {
		t.Errorf("Server config mismatch")
	}
	
	// Verify that the SSH config is set up
	if server.sshConfig == nil {
		t.Errorf("SSH config is nil")
	}
	
	// Verify that the server version is set correctly
	if server.sshConfig.ServerVersion != "SSH-2.0-Test" {
		t.Errorf("Unexpected server version: %s", server.sshConfig.ServerVersion)
	}
}