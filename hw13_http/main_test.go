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
	defer resp.Body.Close()
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
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
	if string(body) != "Response from server" {
		t.Errorf("Unexpected response body: %q", string(body))
	}
}

func TestSendRequestGET(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	body, err := sendRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if body != "Response from server" {
		t.Errorf("Unexpected response body: %q", body)
	}
}

func TestSendRequestPOST(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	requestBody := []byte("test data")
	body, err := sendRequest(http.MethodPost, ts.URL, requestBody)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if body != "Response from server" {
		t.Errorf("Unexpected response body: %q", body)
	}
}
