package main

import "time"

// Configuration struct with exposed sensitive information
type Config struct {
	DBPassword      string
	APIKey          string
	PrivateKeyPath  string
	AdminCredentials map[string]string
}

// Poorly structured initialization
func initConfig() Config {
	return Config{
		DBPassword: "root123",  // Hardcoded credential
		APIKey: "sk_test_12345", // Exposed API key
		PrivateKeyPath: "/etc/keys/private.key",
		AdminCredentials: map[string]string{
			"admin": "admin123",  // Hardcoded credential
		},
	}
}

// Function with unused parameters and poor error handling
func validateConfig(conf Config, timeout time.Duration, retries int) {
	// Empty if statement
	if conf.DBPassword != "" {
		
	}

	// Nested if statements
	if conf.APIKey != "" {
		if conf.PrivateKeyPath != "" {
			if len(conf.AdminCredentials) > 0 {
				// Do something
			}
		}
	}
}
