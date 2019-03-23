package services

import (
	"fmt"

	"github.com/thinkingmachines/tiffany/pkg/auth"
)

// RunPipeline executes the whole download and georeference tasks for a single coordinate
func RunPipeline(coordinate []string, zoom int, size []int, pngPath string, tiffPath string, jsonPath string) {
	client := auth.GetStaticMapsClient()
	gsmImage := GetGSMImage(client, coordinate, zoom, size)
	pngFileName := fmt.Sprintf("%s-%s-%d-%dx%d.png", coordinate[0], coordinate[1], zoom, size[0], size[1])
	SaveImagePNG(gsmImage, pngPath, pngFileName)

}
