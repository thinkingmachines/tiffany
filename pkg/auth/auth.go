package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// GetStaticMapsClient returns a Google Static Maps Client
func GetStaticMapsClient() *maps.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Static Maps Client error: %s", err)
	}

	return client
}
