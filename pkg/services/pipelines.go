// Package services provides functions that enable us to interact with various
// external services such as the Static Maps API and/or Google BigQuery
package services

import (
	"fmt"

	"github.com/thinkingmachines/tiffany/pkg/auth"
)

// RunPipeline executes the whole download and georeference tasks for a single coordinate
func RunPipeline(coordinate []string, zoom int, size []int, path string) {
	client := auth.GetStaticMapsClient()
	gsmImage := GetGSMImage(client, coordinate, zoom, size)
	pngFileName := fmt.Sprintf("%s-%s-%d-%dx%d.png", coordinate[0], coordinate[1], zoom, size[0], size[1])
	SaveImagePNG(gsmImage, fmt.Sprintf("%spng/", path), pngFileName)
	tifFileName := fmt.Sprintf("%s-%s-%d-%dx%d.tif", coordinate[0], coordinate[1], zoom, size[0], size[1])
	GeoreferenceImage(coordinate, size, pngFileName, fmt.Sprintf("%stiff/%s", path, tifFileName))
}
