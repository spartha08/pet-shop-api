package main

import (
	"fmt"
	"os"
)

// Global variable - bad practice
var DEBUG = true

// Function with too many nested if statements
func processRequest(input string) {
	if input != "" {
		if len(input) > 5 {
			if input[0] == 'A' {
				if DEBUG {
					fmt.Println("Processing:", input)
				}
			}
		}
	}
}

// Duplicated code
func validateInput(s string) bool {
	if s == "" {
		return false
	}
	if len(s) < 3 {
		return false
	}
	if len(s) > 50 {
		return false
	}
	return true
}

// Another duplicate
func validateName(s string) bool {
	if s == "" {
		return false
	}
	if len(s) < 3 {
		return false
	}
	if len(s) > 50 {
		return false
	}
	return true
}

// Function with security issues
func checkPermissions(filename string) {
	os.Chmod(filename, 0777) // Security issue - overly permissive file permissions
}

// Poor error handling
func divide(a, b int) int {
	return a / b  // No error handling for division by zero
}
