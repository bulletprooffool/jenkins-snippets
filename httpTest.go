package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Command-line flags
	uri := flag.String("Uri", "", "The URI to test")
	flag.Parse()

	// Check if URI is provided
	if *uri == "" {
		fmt.Println("Please provide a URI using the -Uri flag.")
		os.Exit(1)
	}

	// Disable security checks for simplicity. In a production environment, you should handle this more securely.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	// Make a GET request to the specified URI
	resp, err := client.Get(*uri)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Check if the response status code indicates success (2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("Success! Valid certificate for %s\n", *uri)
		os.Exit(0)
	} else {
		fmt.Printf("Error: Invalid certificate for %s. Status code: %d\n", *uri, resp.StatusCode)
		os.Exit(1)
	}
}
