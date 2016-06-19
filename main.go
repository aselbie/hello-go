package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

type Response struct {
	Body string `json:"body"`
}

func sayHiHandler(w http.ResponseWriter, r *http.Request) {
	var person Person

	// Reject non-POST requests to this endpoint
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Reject request if message body is empty
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&person)

	// Reject request if JSON parsing fails
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Build up missingFields slice
	missingFields := []string{}

	if person.FirstName == "" {
		missingFields = append(missingFields, "firstName")
	}

	if person.LastName == "" {
		missingFields = append(missingFields, "lastName")
	}

	// Reject request if required fields are missing
	if len(missingFields) > 0 {
		http.Error(w, "The following fields were missing from your request: " + strings.Join(missingFields[:],", "), http.StatusBadRequest)
		return
	}

	response := Response{"Hi " + person.FirstName + " " + person.LastName}

	json.NewEncoder(w).Encode(response)
}

// Kickoff server
// Default behavior returns 404 for unhandled routes
func main() {
	http.HandleFunc("/sayhi", sayHiHandler)
	http.ListenAndServe(":3000", nil)
}
