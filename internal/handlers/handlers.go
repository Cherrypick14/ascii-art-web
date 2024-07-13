package handlers

import (
    "net/http"
    "html/template"
    "asciiweb/internal/ascii"
    "log"
)

var (
    templates = template.Must(template.ParseFiles("internal/templates/index.html"))
)

type PageData struct {
    Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
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

    if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
