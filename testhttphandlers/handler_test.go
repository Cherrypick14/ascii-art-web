package testhttphandlers

import (
	"net/http"
	"net/http/httptest"
     "testing"
	"asciiweb/handlers"
	"strings"
)

type HomeHandlerTestCase struct {
	method         string
	target         string
	expectedStatus int
}

func TestHomeHandler(t *testing.T) {
	testCases := []HomeHandlerTestCase{
		{
			method:         "GET",
			target:         "/",
			expectedStatus: http.StatusOK,
		},
		{
			method:         "POST",
			target:         "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			method:         "PUT",
			target:         "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			method:         "GET",
			target:         "/non-existent",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.target, nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.HomeHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != tc.expectedStatus {
			t.Errorf("Expected status code %v, but got %v", tc.expectedStatus, status)
		}

	}
}

type AsciiArtHandlerTestCase struct {
	method         string
	target         string
	formData       string
	expectedStatus int
}

func TestAsciiArtHandler(t *testing.T) {
	testCases := []AsciiArtHandlerTestCase{
		{
			method:         "POST",
			target:         "/ascii",
			formData:       "text=Hello%2C+world!&banner=standard",
			expectedStatus: http.StatusOK,
		},
		{
			method:         "GET",
			target:         "/ascii",
			formData:       "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			method:         "POST",
			target:         "/ascii",
			formData:       "text=Hello%2C+world!",
			expectedStatus: http.StatusBadRequest,
		},
		{
			method:         "POST",
			target:         "/ascii",
			formData:       "banner=standard",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.target, strings.NewReader(tc.formData))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.AsciiArtHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != tc.expectedStatus {
			t.Errorf("Expected status code %v, but got %v", tc.expectedStatus, status)
		}
	}
}
