package analysis

import (
	"time"
)

// AttackData represents data about an authentication attempt
type AttackData struct {
	Timestamp time.Time
	IP        string
	Username  string
	Method    string
	Password  string // For password auth
}

// Analyzer analyzes attack patterns
type Analyzer struct {
	attacks []AttackData
}

// NewAnalyzer creates a new analyzer
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		attacks: make([]AttackData, 0),
	}
}

// AddAttack adds an attack to the analyzer
func (a *Analyzer) AddAttack(ip, username, method, credential string) {
	attack := AttackData{
		Timestamp: time.Now(),
		IP:        ip,
		Username:  username,
		Method:    method,
	}
	
	if method == "password" {
		attack.Password = credential
	}
	
	a.attacks = append(a.attacks, attack)
}

// GetTopUsernames returns the most commonly attempted usernames
func (a *Analyzer) GetTopUsernames(limit int) map[string]int {
	counts := make(map[string]int)
	
	for _, attack := range a.attacks {
		counts[attack.Username]++
	}
	
	// This is a stub - in a real implementation, we would sort and return top N
	return counts
}

// GetTopPasswords returns the most commonly attempted passwords
func (a *Analyzer) GetTopPasswords(limit int) map[string]int {
	counts := make(map[string]int)
	
	for _, attack := range a.attacks {
		if attack.Password != "" {
			counts[attack.Password]++
		}
	}
	
	// This is a stub - in a real implementation, we would sort and return top N
	return counts
}

// GetTopIPs returns the most common attacker IPs
func (a *Analyzer) GetTopIPs(limit int) map[string]int {
	counts := make(map[string]int)
	
	for _, attack := range a.attacks {
		counts[attack.IP]++
	}
	
	// This is a stub - in a real implementation, we would sort and return top N
	return counts
}