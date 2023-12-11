package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ldap.v3"
)

func main() {
	// Command-line flags
	uri := flag.String("Uri", "", "The URI to test (e.g., ldaps://example.com:636)")
	flag.Parse()

	// Check if URI is provided
	if *uri == "" {
		fmt.Println("Please provide a URI using the -Uri flag.")
		os.Exit(1)
	}

	// Parse URI to get host and port
	parts := strings.Split(*uri, "://")
	if len(parts) != 2 {
		fmt.Println("Invalid URI format. Please use a valid URI (e.g., ldaps://example.com:636).")
		os.Exit(1)
	}

	host := parts[1]
	if !strings.Contains(host, ":") {
		fmt.Println("Please specify a port in the URI (e.g., ldaps://example.com:636).")
		os.Exit(1)
	}

	// Disable security checks for simplicity. In a production environment, you should handle this more securely.
	ldap.DefaultTimeout = 10 // seconds
	l, err := ldap.DialTLS("tcp", host, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()

	// Make a simple bind to check if the connection is successful
	err = l.Bind("", "")
	if err != nil {
		fmt.Printf("Error: Unable to bind. %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Success! Connected to %s\n", *uri)
}
