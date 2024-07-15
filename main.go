package main

import (
	"asciiweb/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root path by serving the home page
	http.HandleFunc("/", handlers.HomeHandler)
	// Handle the /ascii-art path by processing the form submission
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	// Start the server on port 8080
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
