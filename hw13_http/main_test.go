package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
	if string(body) != "Response from server" {
		t.Errorf("Unexpected response body: %q", string(body))
	}
}

func TestHandlerPOST(t *testing.T) {
	requestBody := "test data"
	req := httptest.NewRequest(http.MethodPost, "http://localhost/", bytes.NewBufferString(requestBody))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
	if string(body) != "Response from server" {
		t.Errorf("Unexpected response body: %q", string(body))
	}
}
