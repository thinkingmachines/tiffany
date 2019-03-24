// Package services provides functions that enable us to interact with various
// external services such as the Static Maps API and/or Google BigQuery
package services

import (
	"fmt"
	"path/filepath"

	"github.com/thinkingmachines/tiffany/pkg/auth"
)

// RunPipeline executes the whole download and georeference tasks for a single coordinate
func RunPipeline(coordinate []string, zoom int, size []int, path string) {

	const gsmSubDir string = "png"
	const geoSubDir string = "tif"

	// Create filenames for output artifacts
	fnameFormat := fmt.Sprintf("%s-%s-%d-%dx%d", coordinate[0], coordinate[1], zoom, size[0], size[1])
	pngPath := filepath.Join(path, gsmSubDir, fnameFormat+".png")
	tifPath := filepath.Join(path, geoSubDir, fnameFormat+".tiff")

	client := auth.GetStaticMapsClient()
	gsmImage := GetGSMImage(client, coordinate, zoom, size)
	SaveImagePNG(gsmImage, pngPath)
	GeoreferenceImage(coordinate, size, pngPath, tifPath)
}
