package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/capture", captureHandler)
	http.HandleFunc("/debit", debitHandler)
	http.HandleFunc("/credit", creditHandler)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func captureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)

	time.Sleep(3 * time.Second)
	if randomNumber < 1 {
		time.Sleep(3 * time.Second)
		http.Error(w, "Error occurred during capture", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Capture method called")
}

func debitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Debit method called")
}

func creditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Credit method called")
}
