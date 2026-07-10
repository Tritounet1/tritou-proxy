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
	cfg := loadConfig()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	fmt.Printf("%+v\n", cfg)

	targets := buildTargets(cfg)

	proxy := buildProxy(debug, targets)

	mux.HandleFunc("/health", health)
	mux.Handle("/", routeHandler(&proxy, targets))

	fmt.Println("Server starting on port", port)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
