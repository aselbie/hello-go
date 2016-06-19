package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

type Response struct {
	Body string `json:"body"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var person Person

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(person)

	response := Response{"Hi " + person.FirstName + " " + person.LastName}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
