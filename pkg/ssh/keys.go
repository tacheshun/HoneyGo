package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// generateTemporaryKey generates a temporary SSH key
func generateTemporaryKey(keyType string) (ssh.Signer, error) {
	switch keyType {
	case "rsa":
		return generateRSAKey()
	default:
		return generateRSAKey()
	}
}

// generateRSAKey generates a temporary RSA key
func generateRSAKey() (ssh.Signer, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}

	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer from key: %w", err)
	}

	return signer, nil
}