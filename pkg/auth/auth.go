// Package auth provides helper functions for authenticating with external
// services such as the Google Static Maps and Google Cloud Platform.
package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// GetStaticMapsClient returns a Client for constructing a StaticMapRequest.
func GetStaticMapsClient() *maps.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_KEY")

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Maps Client error: %s", err)
	}

	return client
}
