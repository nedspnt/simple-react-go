package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// Create a new CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	// Message Endpoint
	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		handleMessage(w, r)
	})

	// Other Endpoint

	// Wrap your handler with the CORS middleware
	handler := corsMiddleware.Handler(http.DefaultServeMux)

	// Start the server on port 8080
	http.ListenAndServe(":8080", handler)

}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	// create a simple message
	message := Message{Text: "Hello from Golang backend!"}

	// convert the message to JSON
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
