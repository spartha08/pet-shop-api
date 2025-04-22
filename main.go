package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Global variables - bad practice
var (
	pets        []Pet
	PASSWORD    = "admin123"  // Hardcoded credentials - security issue
	SECRET_KEY  = "mysecret"  // Hardcoded secret - security issue
	db_password = "root123"   // Naming convention violation
)

// Pet struct with exported fields but no documentation
type Pet struct {
	Id       int
	Name     string
	Type     string
	age      int    // Unexported field - inconsistent
	owner    string // Unexported field - inconsistent
	Password string // Sensitive data exposure
}

// Duplicate code in multiple functions
func duplicateFunction1() {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
		fmt.Println(result)
		fmt.Println(result)
		fmt.Println(result)
	}
}

func duplicateFunction2() {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
		fmt.Println(result)
		fmt.Println(result)
		fmt.Println(result)
	}
}

// Function with too many parameters
func createPetWithDetails(id int, name string, type_ string, age int, owner string, color string, weight float64, height float64, length float64, breed string) {
	// Function body
}

// SQL Injection vulnerability
func getPetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "SELECT * FROM pets WHERE id = " + id // SQL Injection vulnerability
	fmt.Fprintf(w, "Executing query: %s", query)
}

// Resource leak - file not properly closed
func writeLog(message string) {
	f, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(message)
	// File never closed - resource leak
}

// Main handler function with multiple issues
func handleRequests() {
	http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(pets)
		case "POST":
			var pet Pet
			err := json.NewDecoder(r.Body).Decode(&pet)
			if err != nil {
				// Error handling with sensitive information exposure
				http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
				return
			}
			
			// No input validation
			pets = append(pets, pet)
			w.WriteHeader(http.StatusCreated)
		}
	})

	// Handling pet by ID with multiple issues
	http.HandleFunc("/pets/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/pets/"))
		
		// Inefficient loop
		for i := 0; i < len(pets); i++ {
			if pets[i].Id == id {
				if r.Method == "GET" {
					json.NewEncoder(w).Encode(pets[i])
					return
				}
				if r.Method == "DELETE" {
					// Unsafe concurrent access
					pets = append(pets[:i], pets[i+1:]...)
					return
				}
			}
		}
		
		// Status code should be 404
		w.WriteHeader(http.StatusOK)
	})

	// Authentication with security issues
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		
		if password == PASSWORD {  // Hardcoded password comparison
			fmt.Fprintf(w, "Welcome %s!", username)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Initialization with magic numbers
	pets = make([]Pet, 0, 100)
	
	// Debug information left in code
	fmt.Println("Debug: Starting server...")
	fmt.Printf("Debug: Using password: %s\n", PASSWORD)
	
	handleRequests()
}
