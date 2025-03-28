# HoneyGO - WORK IN PROGRESS
HoneyGO is a lightweight, low-interaction SSH honeypot written in Go. It simulates an SSH server to collect and analyze brute force attacks and unauthorized access attempts.

## Key Features

* Simulates SSH login prompts and logs credentials
* Captures attack patterns and common brute-force sources
* Configurable server banner to mimic different SSH implementations
* Detailed logging of all connection attempts and authentication data
* Basic analysis of attack patterns
* Optionally forwards attackers to a real honeypot (e.g., Cowrie)
  
## Security Considerations

* Run HoneyGO with the principle of least privilege
* Consider using a separate user account for the honeypot
* Do not run the honeypot on production systems
* Regularly review and analyze the collected data

## WORK IN PROGRESS 
