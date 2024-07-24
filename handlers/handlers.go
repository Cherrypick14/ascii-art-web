package handlers

import (
	"html/template"
	"log"
	"net/http"

	"asciiweb/ascii"
)

type PageData struct {
	Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		log.Println("Resource not found")
		return
	}
	// Render the form
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "404 Page Not Found: Error parsing template", http.StatusNotFound)
		log.Println("template/index.html not found")
		return
	}
	// Create an instance of PageData with the desired result
	data := PageData{
		Result: "ASCII Art Generator", // Replace with your actual result(Ascii-art)
	}

	// Execute the template with the provided data
	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("I'm out")
		return
	}

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404: Page Not found", http.StatusNotFound)
		log.Println("404: Page Not found")

	} else {
		// Process the form and generate ASCII art
		input := r.FormValue("text")
		banner := r.FormValue("banner")

		if input == "" || banner == "" {
			http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
			log.Println("400 Bad Request: Input or Banner is empty")

			return
		}

		asciiArt, err := ascii.AsciiArt(input, banner)
		if err != nil {
			http.Error(w, "Unsupported input: Cannot generate ASCII art for the provided text", http.StatusBadRequest)
			log.Printf("400 status code: %v", err)

			return
		}

		data := PageData{
			Result: asciiArt,
		}

		w.WriteHeader(http.StatusOK)

		log.Println("200 OK: ASCII art generated successfully")

		// Render the template with the generated ASCII art
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("500 Internal Server Error: Error executing template: %v", err)
			return
		}

		tpl.Execute(w, data)
	}
}
