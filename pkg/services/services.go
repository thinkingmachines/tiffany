// Package services provides functions that enable us to interact with various
// external services such as the Static Maps API and/or Google BigQuery
package services

import (
	"context"
	"fmt"
	"image"
	"log"
	"thinkingmachines/tiffany/pkg/types"
)

// GetGSMImage downloads a single static maps image given a client and set of
// parameters
func GetGSMImage(client *maps.Client, coordinate types.Coordinate, zoom int, size types.Size, maptype string) image.Image {
	// Prepare request
	r := &maps.StaticMapRequest{
		Center:  fmt.Sprintf("%s,%s", coordinate.Latitude, coordinate.Longitude),
		Zoom:    zoom,
		Size:    fmt.Sprintf("%sx%s", size.Length, size.Width),
		Scale:   2,
		MapType: "satellite",
	}
	// Perform request
	img, err := client.StaticMap(context.Background(), r)
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	return img
}
