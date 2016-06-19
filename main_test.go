package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fullJsonString = []byte(`{"firstName": "Jon", "lastName": "Snow"}`)
var missingJsonString = []byte(`{"firstName": "Jon"}`)
var invalidJsonString = []byte(`{"firstName": "Jon", "last`)

func TestValidRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(fullJsonString))
	w := httptest.NewRecorder()

	sayHiHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("A valid request should return %v", http.StatusOK)
	}
}

func TestInvalidMethod(t *testing.T) {
	req, _ := http.NewRequest("GET", "", bytes.NewBuffer(fullJsonString))
	w := httptest.NewRecorder()

	sayHiHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Methods other than POST should return %v", http.StatusMethodNotAllowed)
	}
}

func TestNoBody(t *testing.T) {
	req, _ := http.NewRequest("POST", "", nil)
	w := httptest.NewRecorder()

	sayHiHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Requests without a body should return %v", http.StatusBadRequest)
	}
}

func TestMissingField(t *testing.T) {
	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(missingJsonString))
	w := httptest.NewRecorder()

	sayHiHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Requests with missing fields should return %v", http.StatusBadRequest)
	}
}

func TestInvalidJson(t *testing.T) {
	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(invalidJsonString))
	w := httptest.NewRecorder()

	sayHiHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Requests without a body should return %v", http.StatusBadRequest)
	}
}
