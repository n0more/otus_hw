package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/spf13/pflag"
)

func parseArguments() (string, []string) {
	var mode string
	pflag.StringVarP(&mode, "mode", "m", "", "Mode (client|server)")
	pflag.Parse()
	return mode, pflag.Args()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	defer r.Body.Close()

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Request body: %s\n", string(body))
	}
	if r.Method != "POST" && r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Response from server"))
}

func sendRequest(method, url string, data []byte) (string, error) {
	var req *http.Request
	var err error

	switch method {
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, url, nil)
	default:
		return "", fmt.Errorf("invalid method: %s", method)
	}

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	mode, aArg := parseArguments()

	switch mode {
	case "client":
		if len(aArg) < 2 {
			fmt.Println("Usage: --mode client <fullURL> <method> [data]")
			return
		}

		fullURL := aArg[0]
		method := aArg[1]
		var data []byte
		if method == http.MethodPost && len(aArg) > 2 {
			data = []byte(aArg[2])
		}

		response, err := sendRequest(method, fullURL, data)
		if err != nil {
			log.Println("Error sending request:", err)
			return
		}
		fmt.Println("Response:", response)
	case "server":
		if len(aArg) != 2 {
			fmt.Println("Usage: -m server <address> <port>")
			return
		}

		address := aArg[0]
		port := aArg[1]
		serverAddr := net.JoinHostPort(address, port)

		http.HandleFunc("/", handler)

		fmt.Printf("Starting server at %s\n", serverAddr)
		server := &http.Server{
			Addr:         serverAddr,
			Handler:      nil,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Println("Server failed:", err)
		}
	default:
		fmt.Println("Invalid mode. Use -m client or -m server")
		return
	}
}
