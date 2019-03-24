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

// GeoreferenceImage converts a Static Maps image into a geo-referenced TIFF
func GeoreferenceImage(coordinate []string, size []int, inpath string, outpath string) {

	log.Println("Georeferencing image")
	// Define projection constants
	const epsg int = 3857
	const projector float64 = 156543.03392
	const maxExtent float64 = 20037508.34

	lat4326, err := strconv.ParseFloat(coordinate[0], 64)
	if err != nil {
		log.Panic(err)
	}
	lon4326, err := strconv.ParseFloat(coordinate[1], 64)
	if err != nil {
		log.Panic(err)
	}

	latCenter := size[0] / 2
	lonCenter := size[1] / 2

	// Compute the GSD Resolution and convert to EPSG3857
	gsdResolution := projector * math.Cos(lat4326*math.Pi/180) / math.Pow(2, 17)
	lat3857 := math.Log((math.Tan(90+lat4326) * math.Pi / 360)) / (math.Pi / 180) * maxExtent / 180
	lon3857 := lon4326 * maxExtent / 180

	// Compute boundaries
	upperLeftY := lat3857 + gsdResolution*float64(latCenter)
	upperLeftX := lon3857 - gsdResolution*float64(lonCenter)

	// Specify raster location through the geotransform array
	gt := [6]float64{upperLeftX, gsdResolution, 0, upperLeftY, 0, -gsdResolution}

	// Read source image and its driver
	srcDataset, err := gdal.Open(inpath, gdal.ReadOnly)
	if err != nil {
		log.Panic(err)
	}
	driver, err := gdal.GetDriverByName("GTiff")
	if err != nil {
		log.Panic(err)
	}

	// Create copy for destination image and update it
	dstDataset := driver.CreateCopy(outpath, srcDataset, 0, nil, nil, nil)
	dstDataset.SetGeoTransform(gt)

	// Get raster projection
	spatialRef := gdal.CreateSpatialReference("")
	spatialRef.FromEPSG(epsg)
	srString, err := spatialRef.ToWKT()
	dstDataset.SetProjection(srString)
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
