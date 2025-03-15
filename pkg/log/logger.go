package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Logger handles logging for the honeypot
type Logger struct {
	logFile *os.File
	mu      sync.Mutex
}

// NewLogger creates a new logger
func NewLogger(path string) (*Logger, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open log file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &Logger{
		logFile: file,
	}, nil
}

// Close closes the logger
func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	if l.logFile != nil {
		l.logFile.Close()
		l.logFile = nil
	}
}

// Info logs informational messages
func (l *Logger) Info(format string, args ...interface{}) {
	l.log("INFO", format, args...)
}

// Error logs error messages
func (l *Logger) Error(format string, args ...interface{}) {
	l.log("ERROR", format, args...)
}

// Auth logs authentication attempts
func (l *Logger) Auth(ip, username, method, credential string) {
	l.log("AUTH", "Login attempt from %s with username '%s' using %s authentication: %s", 
		ip, username, method, credential)
	
	// Also log to standard output for immediate visibility
	fmt.Printf("[%s] AUTH: Login attempt from %s with username '%s' using %s authentication\n",
		time.Now().Format("2006-01-02 15:04:05"), ip, username, method)
}

// log writes a log message to the log file
func (l *Logger) log(level, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	if l.logFile == nil {
		return
	}
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	
	logLine := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
	
	_, err := l.logFile.WriteString(logLine)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log file: %v\n", err)
	}
}