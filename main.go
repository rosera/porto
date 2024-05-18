package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
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
        response := service.Default

        // Create a closure to capture the response.
        http.HandleFunc("/"+endpoint, func(response string) http.HandlerFunc {
            return func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "<h1>%s</h1>", response)
            }
        }(response))
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
