# Ascii-Art-Web

## Description
Ascii-art-web is a web application that allows users to generate ASCII art using different banners. It features a graphical user interface for inputting text and selecting a banner style.

## Authors
- Cherrypick14

## Usage
1. Run the server: `go run main.go`
2. Open your browser and navigate to `http://localhost:8080`

## Implementation details
- The server is implemented in Go.
- ASCII art generation logic is in `ascii/ascii.go`.
- HTTP handlers are in `handlers/handlers.go`.
- HTML templates are in `templates/`.

## Instructions
1. Clone the repository.
2. Place the banner files (`shadow.txt`, `standard.txt`, `thinkertoy.txt`) in `ascii/`.
3. Run the server: `go run main.go`
4. Access the web application at `http://localhost:8080`.

## HTTP Status Codes
- `200 OK`: Successful response.
- `404 Not Found`: Resource not found.
- `400 Bad Request`: Incorrect request.
- `500 Internal Server Error`: Unhandled server error.
