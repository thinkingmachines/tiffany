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
(c) Thinking Machines Data Science, 2019`,
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
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		coordinate := []string{args[0], args[1]}
		RunPipeline(coordinate, zoom, size, path, noRef, wtLbl)
	},
}

var zoom int
var size []int
var path string
var wtLbl string
var noRef bool

func init() {
	rootCmd.PersistentFlags().IntVarP(&zoom, "zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceVarP(&size, "size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().StringVar(&path, "path", "tiffany.out/", "path to save output artifacts")
	rootCmd.PersistentFlags().StringVar(&wtLbl, "with-labels", "", "path to the label's WKT representation (.csv)")
	rootCmd.PersistentFlags().BoolVar(&noRef, "without-reference", false, "do not georeference")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
