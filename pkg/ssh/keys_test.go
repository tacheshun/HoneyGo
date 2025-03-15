package ssh

import (
	"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	// Generate a key
	key, err := generateRSAKey()
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}
	
	// Verify that it's a valid key
	if key == nil {
		t.Fatalf("Generated key is nil")
	}
	
	// Check that we can get the public key
	pubKey := key.PublicKey()
	if pubKey == nil {
		t.Fatalf("Public key is nil")
	}
	
	// Verify the key type is RSA
	keyType := pubKey.Type()
	if keyType != "ssh-rsa" {
		t.Errorf("Expected key type to be ssh-rsa, got %s", keyType)
	}
}

func TestGenerateTemporaryKey(t *testing.T) {
	// Test with RSA key type
	key, err := generateTemporaryKey("rsa")
	if err != nil {
		t.Fatalf("Failed to generate temporary RSA key: %v", err)
	}
	
	if key == nil {
		t.Fatalf("Generated RSA key is nil")
	}
	
	if key.PublicKey().Type() != "ssh-rsa" {
		t.Errorf("Expected key type to be ssh-rsa, got %s", key.PublicKey().Type())
	}
	
	// Test with unsupported key type (should default to RSA)
	key, err = generateTemporaryKey("unsupported")
	if err != nil {
		t.Fatalf("Failed to generate default key for unsupported type: %v", err)
	}
	
	if key == nil {
		t.Fatalf("Generated default key is nil")
	}
	
	if key.PublicKey().Type() != "ssh-rsa" {
		t.Errorf("Expected default key type to be ssh-rsa, got %s", key.PublicKey().Type())
	}
}