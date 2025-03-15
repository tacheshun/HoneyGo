package analysis

import (
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	analyzer := NewAnalyzer()
	
	if analyzer == nil {
		t.Fatalf("NewAnalyzer returned nil")
	}
	
	if analyzer.attacks == nil {
		t.Errorf("Attacks slice should be initialized, but is nil")
	}
	
	if len(analyzer.attacks) != 0 {
		t.Errorf("Expected empty attacks slice, got %d elements", len(analyzer.attacks))
	}
}

func TestAddAttack(t *testing.T) {
	analyzer := NewAnalyzer()
	
	// Add a password attack
	analyzer.AddAttack("192.168.1.1", "admin", "password", "password123")
	
	if len(analyzer.attacks) != 1 {
		t.Fatalf("Expected 1 attack, got %d", len(analyzer.attacks))
	}
	
	attack := analyzer.attacks[0]
	
	if attack.IP != "192.168.1.1" {
		t.Errorf("Expected IP to be 192.168.1.1, got %s", attack.IP)
	}
	
	if attack.Username != "admin" {
		t.Errorf("Expected username to be admin, got %s", attack.Username)
	}
	
	if attack.Method != "password" {
		t.Errorf("Expected method to be password, got %s", attack.Method)
	}
	
	if attack.Password != "password123" {
		t.Errorf("Expected password to be password123, got %s", attack.Password)
	}
	
	// Add a key-based attack
	analyzer.AddAttack("192.168.1.2", "root", "publickey", "key-data")
	
	if len(analyzer.attacks) != 2 {
		t.Fatalf("Expected 2 attacks, got %d", len(analyzer.attacks))
	}
	
	attack = analyzer.attacks[1]
	
	if attack.IP != "192.168.1.2" {
		t.Errorf("Expected IP to be 192.168.1.2, got %s", attack.IP)
	}
	
	if attack.Username != "root" {
		t.Errorf("Expected username to be root, got %s", attack.Username)
	}
	
	if attack.Method != "publickey" {
		t.Errorf("Expected method to be publickey, got %s", attack.Method)
	}
	
	// For non-password auth methods, Password should be empty
	if attack.Password != "" {
		t.Errorf("Expected password to be empty for key-based auth, got %s", attack.Password)
	}
}

func TestAnalysis(t *testing.T) {
	analyzer := NewAnalyzer()
	
	// Add some test data
	analyzer.AddAttack("192.168.1.1", "admin", "password", "password123")
	analyzer.AddAttack("192.168.1.1", "root", "password", "password123")
	analyzer.AddAttack("192.168.1.2", "admin", "password", "admin123")
	analyzer.AddAttack("192.168.1.3", "ubuntu", "password", "ubuntu")
	analyzer.AddAttack("192.168.1.3", "ubuntu", "password", "ubuntu")
	
	// Test GetTopUsernames
	usernames := analyzer.GetTopUsernames(10)
	
	if usernames["admin"] != 2 {
		t.Errorf("Expected 'admin' to have count 2, got %d", usernames["admin"])
	}
	
	if usernames["root"] != 1 {
		t.Errorf("Expected 'root' to have count 1, got %d", usernames["root"])
	}
	
	if usernames["ubuntu"] != 2 {
		t.Errorf("Expected 'ubuntu' to have count 2, got %d", usernames["ubuntu"])
	}
	
	// Test GetTopPasswords
	passwords := analyzer.GetTopPasswords(10)
	
	if passwords["password123"] != 2 {
		t.Errorf("Expected 'password123' to have count 2, got %d", passwords["password123"])
	}
	
	if passwords["admin123"] != 1 {
		t.Errorf("Expected 'admin123' to have count 1, got %d", passwords["admin123"])
	}
	
	if passwords["ubuntu"] != 2 {
		t.Errorf("Expected 'ubuntu' to have count 2, got %d", passwords["ubuntu"])
	}
	
	// Test GetTopIPs
	ips := analyzer.GetTopIPs(10)
	
	if ips["192.168.1.1"] != 2 {
		t.Errorf("Expected '192.168.1.1' to have count 2, got %d", ips["192.168.1.1"])
	}
	
	if ips["192.168.1.2"] != 1 {
		t.Errorf("Expected '192.168.1.2' to have count 1, got %d", ips["192.168.1.2"])
	}
	
	if ips["192.168.1.3"] != 2 {
		t.Errorf("Expected '192.168.1.3' to have count 2, got %d", ips["192.168.1.3"])
	}
}