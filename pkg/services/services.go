// Package services provides functions that enable us to interact with various
// external services such as the Static Maps API and/or Google BigQuery
package services

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/thinkingmachines/tiffany/pkg/auth"
	"googlemaps.github.io/maps"
)

// GetGSMImage downloads a single static maps image given a client and set of
// parameters
func GetGSMImage(client *maps.Client, coordinate []string, zoom int, size []int) image.Image {
	// Prepare request
	r := &maps.StaticMapRequest{
		Center:  fmt.Sprintf("%s,%s", coordinate[0], coordinate[1]),
		Zoom:    zoom,
		Size:    fmt.Sprintf("%dx%d", size[0], size[1]),
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

// SaveImagePNG exports an image into the given image type (PNG or TIFF)
func SaveImagePNG(img image.Image, path string, fname string) {
	log.Printf("Saving image to %s", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	f, err := os.Create(fmt.Sprintf("%s%s", path, fname))
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()
	png.Encode(f, img)
}

// RunPipeline executes the whole download and georeference tasks for a single coordinate
func RunPipeline(coordinate []string, zoom int, size []int, pngPath string, tiffPath string, jsonPath string) {
	client := auth.GetStaticMapsClient()
	gsmImage := GetGSMImage(client, coordinate, zoom, size)
	pngFileName := fmt.Sprintf("%s-%s-%d-%dx%d.png", coordinate[0], coordinate[1], zoom, size[0], size[1])
	SaveImagePNG(gsmImage, pngPath, pngFileName)

}
