// Package services provides functions that enable us to interact with various
// external services such as the Static Maps API and/or Google BigQuery
package services

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lukeroth/gdal"
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
		Scale:   1,
		MapType: "satellite",
	}
	// Perform request
	img, err := client.StaticMap(context.Background(), r)
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	return img
}

// GeoreferenceImage converts a Static Maps image into a geo-referenced TIFF
func GeoreferenceImage(coordinate []string, size []int, inpath string, outpath string) {

	log.Println("Georeferencing image")
	// Define projection constants
	const projector float64 = 156543.03392
	const maxExtent float64 = 20037508.34

	if _, err := os.Stat(filepath.Dir(outpath)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(outpath), os.ModePerm)
	}

	lat4326, _ := strconv.ParseFloat(coordinate[0], 64)
	lon4326, _ := strconv.ParseFloat(coordinate[1], 64)

	latCenter := size[0] / 2
	lonCenter := size[1] / 2

	// Compute the GSD Resolution and convert to EPSG3857
	gsdResolution := projector * math.Cos(lat4326*math.Pi/180) / math.Pow(2, 17)
	lat3857 := (math.Log(math.Tan((90+lat4326)*math.Pi/360)) / (math.Pi / 180)) * maxExtent / 180
	lon3857 := lon4326 * maxExtent / 180

	// Compute boundaries
	upperLeftY := lat3857 + gsdResolution*float64(latCenter)
	upperLeftX := lon3857 - gsdResolution*float64(lonCenter)

	// Read source image and its driver
	srcDataset, _ := gdal.Open(inpath, gdal.ReadOnly)
	driver, _ := gdal.GetDriverByName("GTiff")

	// Open destination dataset
	dstDataset := driver.CreateCopy(outpath, srcDataset, 0, nil, nil, nil)
	dstDataset.SetGeoTransform([6]float64{upperLeftX, gsdResolution, 0, upperLeftY, 0, -gsdResolution})

	// Get raster projection
	srs := gdal.CreateSpatialReference("")
	srs.FromEPSG(3857)
	destWKT, _ := srs.ToWKT()

	// Create destination dataset
	dstDataset.SetProjection(destWKT)

	defer dstDataset.Close()
	defer srcDataset.Close()
}

// SaveImagePNG exports an image into a file
func SaveImagePNG(img image.Image, path string) {
	log.Printf("Saving image to %s", path)
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
	}

	f, err := os.Create(fmt.Sprintf("%s", path))
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()
	png.Encode(f, img)
}
