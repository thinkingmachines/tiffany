package cmd

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

	"github.com/disintegration/imaging"
	"github.com/joho/godotenv"
	"github.com/lukeroth/gdal"
	"googlemaps.github.io/maps"
)

// GeoreferenceImage converts a Static Maps image into a geo-referenced TIFF
func GeoreferenceImage(coordinate []string, size []int, zoom int, inpath string, outpath string) {

	log.Printf("Georeferencing image into %s", outpath)
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
	gsdResolution := projector * math.Cos(lat4326*math.Pi/180) / math.Pow(2, float64(zoom))
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

	dstDataset.SetProjection(destWKT)

	defer dstDataset.Close()
	defer srcDataset.Close()
}

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

// GetStaticMapsClient returns a Client for constructing a StaticMapRequest.
func GetStaticMapsClient() *maps.Client {
	log.Printf("Loading Google Static Maps Client")
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

// ReprojectImage converts the projection of an image back to 4326
func ReprojectImage(path string) {

	options := []string{"-t_srs", "epsg:4326"}
	ds, err := gdal.Open(path, gdal.ReadOnly)
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	out := gdal.GDALWarp(path, gdal.Dataset{}, []gdal.Dataset{ds}, options)
	defer out.Close()
}

// RunPipeline executes the whole download and georeference tasks for a single coordinate
func RunPipeline(coordinate []string, zoom int, size []int, path string, noRef bool) {

	const gsmSubDir string = "png"
	const geoSubDir string = "tif"

	// Create filenames for output artifacts
	fnameFormat := fmt.Sprintf("%s_%s_%d_%dx%d", coordinate[0], coordinate[1], zoom, size[0], size[1])
	pngPath := filepath.Join(path, gsmSubDir, fnameFormat+".png")
	tifPath := filepath.Join(path, geoSubDir, fnameFormat+".tiff")

	client := GetStaticMapsClient()
	gsmImage := GetGSMImage(client, coordinate, zoom, size)
	SaveImagePNG(gsmImage, pngPath)
	if !noRef {
		GeoreferenceImage(coordinate, size, zoom, pngPath, tifPath)
		ReprojectImage(tifPath)
	}
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
	imgRGBA := imaging.Clone(img)

	defer f.Close()
	png.Encode(f, imgRGBA)
}
