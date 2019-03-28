// Package cmd contains all the helper functions, handlers, and command-line methods
// for building the tiffany command-line interface.
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiffany LATITUDE LONGITUDE",
	Short: "tiffany is a tool for rendering to TIFF any image from Google Static Maps",
	Long: `
	     _   _  __  __
	    | | (_)/ _|/ _|
	    | |_ _| |_| |_ __ _ _ __  _   _
	    | __| |  _|  _/ _' | '_ \| | | |
	    | |_| | | | || (_| | | | | |_| |
	     \__|_|_| |_| \__,_|_| |_|\__, |
				       __/ |
				      |___/

    Render to TIFF any Google Static Maps (GSM) image
       (c) Thinking Machines Data Science, 2019
		  Version: v1.0.0-alpha`,
	Example: `
  tiffany 14.54694524 121.0197543253
  tiffany 14.54694524 121.0197543253 --without-reference
  tiffany 14.54694524 121.0197543253 --with-labels=/path/to/file.shp
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Please input the coordinates: LATITUDE LONGITUDE")
		}
		if len(size) != 2 {
			return errors.New("Requires a size in the form {L},{W}")
		}
		if len(wtLbl) > 0 && noRef {
			return errors.New("Conflicting arguments, cannot make labels without georeferencing the image")
		}
		return nil
	},
	Version: "v1.0.0-alpha",
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		coordinate := []string{args[0], args[1]}
		RunPipeline(coordinate, zoom, size, path, noRef, wtLbl)
	},
}

var batchCmd = &cobra.Command{
	Use:   "batch PATH/TO/FILE.CSV",
	Short: "Apply tiffany on a CSV file of coordinates",
	Long: `
The batch command is a more efficient alternative when running tiffany
on a list of lat-lon coordinates. Instead of using a for-loop, you can
just provide the path to the CSV file, and apply the same parameters as
if you're running tiffany on a single point.

Assumes that the first column is the latitude and the second column is the
longitude.
`,
	Example: `
  tiffany batch coordinates.csv
  tiffany batch coordinates.csv --without-reference
  tiffany batch coordinates.csv --with-labels=/path/to/file.shp
`,
	Args:    cobra.ExactArgs(1),
	Version: "v1.0.0-alpha",
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		csvFile := args[0]
		RunBatchPipeline(csvFile, skipFirst, zoom, size, path, noRef, wtLbl)
	},
}

var zoom int
var size []int
var path string
var wtLbl string
var noRef bool
var skipFirst bool

func init() {
	// Add sub-commands
	rootCmd.AddCommand(batchCmd)

	// Define flags for ROOT command
	rootCmd.PersistentFlags().IntVarP(&zoom, "zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceVarP(&size, "size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().StringVar(&path, "path", "tiffany.out/", "path to save output artifacts")
	rootCmd.PersistentFlags().StringVar(&wtLbl, "with-labels", "", "path to the label's ESRI Shapefile")
	rootCmd.PersistentFlags().BoolVar(&noRef, "without-reference", false, "do not georeference")

	// Define flags for `batch` command
	batchCmd.Flags().BoolVar(&skipFirst, "skip-header-row", false, "skip header row")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
