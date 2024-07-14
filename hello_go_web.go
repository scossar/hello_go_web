package main

import (
	"fmt"
	"log"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Bar")
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
