package testhttphandlers

import (
	// "asciiweb/ascii"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"asciiweb/handlers"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		formData       map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET request to root",
			method:         http.MethodGet,
			url:            "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "ASCII Art Generator", // Assuming this text is in your template
		},
		{
			name:           "POST request with valid data",
			method:         http.MethodPost,
			url:            "/",
			formData:       map[string]string{"text": "Hello", "banner": "standard"},
			expectedStatus: http.StatusOK,
			expectedBody:   "ASCII art", // Part of the expected result
		},
		{
			name:           "POST request with empty text",
			method:         http.MethodPost,
			url:            "/",
			formData:       map[string]string{"text": "", "banner": "standard"},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Input or Banner is empty",
		},
		{
			name:           "POST request with empty banner",
			method:         http.MethodPost,
			url:            "/",
			formData:       map[string]string{"text": "Hello", "banner": ""},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Input or Banner is empty",
		},
		{
			name:           "GET request to non-root path",
			method:         http.MethodGet,
			url:            "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 Not Found",
		},
		{
			name:           "PUT request (not allowed)",
			method:         http.MethodPut,
			url:            "/",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method Not Allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.formData != nil {
				req.PostForm = make(map[string][]string)
				for key, value := range tt.formData {
					req.PostForm.Set(key, value)
				}
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.HomeHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.NotFoundHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "404 Not Found"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}