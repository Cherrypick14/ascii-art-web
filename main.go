package main

import (
	"log"
	"net/http"

	"asciiweb/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the /ascii-art path by processing the form submission
	http.HandleFunc("/", handlers.AsciiArtHandler)

	http.HandleFunc("/nonexistent", handlers.NotFoundHandler)

	// Start the server on port 8080
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
