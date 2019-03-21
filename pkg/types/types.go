// Package types contains some data types to define important parameters or values
package types

// Coordinate defines the latitude and longitude of a certain point
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// ImageSize defines the pixel-size of a certain image
type ImageSize struct {
	Length int
	Width  int
}
