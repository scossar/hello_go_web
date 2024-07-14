package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFooHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(fooHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", rec.Code)
	}

	// expected := "Hello, World!"
	if expected := "Hello from Foo"; rec.Body.String() != expected {
		t.Fatalf("Expected body to be %q; got %q", expected, rec.Body.String())
	}
}

func TestBarHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/bar", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(barHandler)
	handler.ServeHTTP(rec, req)

	if expected := "Hello from Bar"; rec.Body.String() != expected {
		t.Fatalf("Expected bofy to be %q; got %q", expected, rec.Body.String())
	}
}
