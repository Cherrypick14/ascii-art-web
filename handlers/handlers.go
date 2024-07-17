package handlers

import (
	"asciiweb/ascii"
	"html/template"
	"log"
	"net/http"
)

// var templates = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler called")

	var data PageData

	if r.Method == http.MethodPost {
		// Process the form submission
		input := r.FormValue("text")
		banner := r.FormValue("banner")

		// Check if input or banner is empty
		if input == "" || banner == "" {
			log.Println("400 Bad Request: Input or Banner is empty")
			http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
			return
		}

		// Process input and banner
		text := ascii.Processinput(input)
		banner2, err := ascii.ReadBannerFile(banner)
		if err != nil {
			log.Printf("500 Internal Server Error: Error reading banner file: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		asciiArt, err := ascii.AsciiArt(text, banner2)
		if err != nil {
			log.Printf("500 Internal Server Error: Error generating ASCII art: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data.Result = asciiArt
		log.Println("200 OK: ASCII art generated successfully")
	}

	// Render the template
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("500 Internal Server Error: Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	tpl.Execute(w, data)
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("405 Method Not Allowed: Only POST requests are allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("text")
	banner := r.FormValue("banner")

	// Check if input or banner is empty
	if input == "" || banner == "" {
		log.Println("400 Bad Request: Input or Banner is empty")
		http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
		return
	}

	// Process input and banner
	text := ascii.Processinput(input)
	banner2, err := ascii.ReadBannerFile(banner)
	if err != nil {
		log.Printf("500 Internal Server Error: Error reading banner file: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	asciiArt, err := ascii.AsciiArt(text, banner2)
	if err != nil {
		log.Printf("500 Internal Server Error: Error generating ASCII art: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: asciiArt,
	}
	log.Println("200 OK: ASCII art generated successfully")

	// Render the template with the generated ASCII art
	tpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Printf("500 Internal Server Error: Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("404 Not Found: Resource not found")
	http.Error(w, "404 Not Found", http.StatusNotFound )
	
}
