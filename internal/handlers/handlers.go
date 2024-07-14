package handlers

import (
	"asciiweb/internal/ascii"
	"html/template"
	"log"
	"net/http"
)

var (
	templates = template.Must(template.ParseFiles("internal/templates/index.html"))
)

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
            http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
            return
        }

        // Process input and banner
        text := ascii.Processinput(input)
        banner2, err := ascii.ReadBannerFile(banner)
        if err != nil {
            log.Printf("Error reading banner file: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        asciiArt, err := ascii.AsciiArt(text, banner2)
        if err != nil {
            log.Printf("Error generating ASCII art: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        data.Result = asciiArt
    }

    // Render the template
    if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("text")
	banner := r.FormValue("banner")

	// Check if input or banner is empty
	if input == "" || banner == "" {
		http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
		return
	}

	// Process input and banner
	text := ascii.Processinput(input)
	banner2, err := ascii.ReadBannerFile(banner)
	if err != nil {
		log.Printf("Error reading banner file: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	asciiArt, err := ascii.AsciiArt(text, banner2)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: asciiArt,
	}

	// Render the template with the generated ASCII art
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
