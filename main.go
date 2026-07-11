package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	cfg, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error while trying to load config.yaml file %q", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	targets, err := buildTargets(cfg)
	if err != nil {
		log.Fatalf("error while trying to build targets %q", err)
	}

	proxy := buildProxy(debug, targets)

	mux.HandleFunc("/health", health)
	mux.Handle("/", routeHandler(&proxy, targets))

	fmt.Println("Server starting on port", port)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
