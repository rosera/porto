package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config represents the YAML configuration structure.
type Config struct {
	Services []Service `yaml:"services"`
}

// Service represents a single service configuration.
type Service struct {
	Endpoint string `yaml:"endpoint"`
	Comment  string `yaml:"comment"`
	Default  string `yaml:"default"`
	Format   string `yaml:"format"`
}

func main() {
	// Read the YAML configuration file.
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Unmarshal the YAML data into the Config struct.
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Define a handler for each service endpoint.
	for _, service := range config.Services {
		endpoint := service.Endpoint
		defaultResponse := service.Default
		format := service.Format

		// Create a closure to capture the response details.
		http.HandleFunc("/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
			var response []byte
			var contentType string

			// Check if the default response is a URL.
			if strings.HasPrefix(defaultResponse, "http://") || strings.HasPrefix(defaultResponse, "https://") {
				// Fetch content from the URL.
				resp, err := http.Get(defaultResponse)
				if err != nil {
					http.Error(w, "Failed to fetch content from URL", http.StatusInternalServerError)
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					http.Error(w, "Failed to read content from URL", http.StatusInternalServerError)
					return
				}
				response = body
				contentType = resp.Header.Get("Content-Type")
			} else {
				// Use the default response directly.
				response = []byte(defaultResponse)
				contentType = "text/plain"
				switch format {
				case "html":
					contentType = "text/html"
				case "json":
					contentType = "application/json"
				case "plain":
					contentType = "text/plain"
				case "image":
					contentType = "image/jpeg"
				}
			}

			// Determine the response format based on the specified format or content type.
			switch format {
			case "json":
				w.Header().Set("Content-Type", "application/json")
				jsonResponse, err := json.Marshal(map[string]string{"response": string(response)})
				if err != nil {
					http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
					return
				}
				w.Write(jsonResponse)
			case "html":
				w.Header().Set("Content-Type", "text/html")
				fmt.Fprint(w, string(response))
			case "plain":
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprint(w, string(response))
			case "image":
				w.Header().Set("Content-Type", "image/jpeg")
				w.Write(response)
			default:
				w.Header().Set("Content-Type", contentType)
				w.Write(response)
			}
		})
	}

	// Define a handler for the root endpoint.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running and healthy")
	})

	// Read the port from the environment variable or use a default port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the web server.
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
