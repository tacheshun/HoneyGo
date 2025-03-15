# HoneyGO

HoneyGO is a lightweight, low-interaction SSH honeypot written in Go. It simulates an SSH server to collect and analyze brute force attacks and unauthorized access attempts.

## Key Features

* Simulates SSH login prompts and logs credentials
* Captures attack patterns and common brute-force sources
* Configurable server banner to mimic different SSH implementations
* Detailed logging of all connection attempts and authentication data
* Basic analysis of attack patterns
* Optionally forwards attackers to a real honeypot (e.g., Cowrie)

## Installation

### Prerequisites

* Go 1.16 or later

### Building from Source

```bash
# Clone the repository
git clone https://github.com/marius/honeygo.git
cd honeygo

# Build the binary
go build -o honeygo ./cmd/honeygo
```

## Configuration

Copy the example configuration file and modify it as needed:

```bash
cp config.example.yaml config.yaml
```

### Example Configuration

```yaml
# Server settings
listen_address: "0.0.0.0:2222"
host_key_path: "keys/honeygo_key"
host_key_type: "rsa"
banner: "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.4"

# Authentication settings
allow_password_auth: true
allow_key_auth: false

# Forwarding settings
forward_enabled: false
forward_host: "127.0.0.1"
forward_port: 2223

# Logging settings
log_path: "logs/honeygo.log"
```

## Usage

```bash
# Start HoneyGO with default configuration
./honeygo

# Start with a custom configuration file
./honeygo -config /path/to/config.yaml
```

## Security Considerations

* Run HoneyGO with the principle of least privilege
* Consider using a separate user account for the honeypot
* Do not run the honeypot on production systems
* Regularly review and analyze the collected data

## License

This project is licensed under the MIT License - see the LICENSE file for details.
