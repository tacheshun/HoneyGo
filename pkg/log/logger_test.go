package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "logger_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Create a logger with a path in the temporary directory
	logPath := filepath.Join(tempDir, "test.log")
	logger, err := NewLogger(logPath)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	
	// Check that the logger is not nil
	if logger == nil {
		t.Fatalf("Logger is nil")
	}
	
	// Check that the log file was created
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("Log file was not created")
	}
}

func TestNestedLogPath(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "logger_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Create a nested path
	nestedDir := filepath.Join(tempDir, "nested", "path")
	logPath := filepath.Join(nestedDir, "test.log")
	
	// Create the logger
	logger, err := NewLogger(logPath)
	if err != nil {
		t.Fatalf("Failed to create logger with nested path: %v", err)
	}
	defer logger.Close()
	
	// Check that the nested directories were created
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Errorf("Nested directory was not created")
	}
}

func TestLogMethods(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "logger_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Create a logger
	logPath := filepath.Join(tempDir, "test.log")
	logger, err := NewLogger(logPath)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	
	// Log some messages
	logger.Info("Test info message with parameter: %s", "value")
	logger.Error("Test error message with parameter: %d", 123)
	logger.Auth("127.0.0.1", "testuser", "password", "testpass")
	
	// Close the logger to flush writes
	logger.Close()
	
	// Read the log file
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}
	
	// Check that the messages were logged
	logContent := string(content)
	
	if !strings.Contains(logContent, "INFO: Test info message with parameter: value") {
		t.Errorf("Info message not found in log file")
	}
	
	if !strings.Contains(logContent, "ERROR: Test error message with parameter: 123") {
		t.Errorf("Error message not found in log file")
	}
	
	if !strings.Contains(logContent, "AUTH: Login attempt from 127.0.0.1 with username 'testuser' using password authentication: testpass") {
		t.Errorf("Auth message not found in log file")
	}
}