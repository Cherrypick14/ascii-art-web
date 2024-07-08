package handlers

import (
	"net/http"
	"html/template"
	"asciiweb/internal/ascii"
	"fmt"
	"log"

)

type PageData struct {
    Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	if err := t.Execute(w, nil); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    input := r.FormValue("text")
    banner := r.FormValue("banner")

	text := ascii.Processinput(input)
	banner2, err := ascii.ReadBannerFile(banner)

	if err != nil {
		fmt.Print("Error reading banner file")
	}
    

    asciiArt, err := ascii.AsciiArt(text, banner2)
	log.Printf("Error generating ASCII art: %v", err)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := PageData{
        Result: asciiArt,
    }

    t, err := template.ParseFiles("templates/index.html")
	log.Printf("Error parsing template: %v", err)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, data); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}