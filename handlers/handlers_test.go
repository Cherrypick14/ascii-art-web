package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		method         string
		url            string
		formData       map[string]string
		expectedStatus int
		expectedBody   string
	}{
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
			var req *http.Request
			var err error

			if tt.method == http.MethodPost {
				formData := url.Values{}
				for key, value := range tt.formData {
					formData.Set(key, value)
				}
				req, err = http.NewRequest(tt.method, tt.url, strings.NewReader(formData.Encode()))
				if err != nil {
					t.Fatalf("could not create request: %v", err)
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(tt.method, tt.url, nil)
				if err != nil {
					t.Fatalf("could not create request: %v", err)
				}
			}

			rr := httptest.NewRecorder()

			// Wrap HomeHandler with the args struct
			args := args{
				w: rr,
				r: req,
			}

			HomeHandler(args.w, args.r)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
