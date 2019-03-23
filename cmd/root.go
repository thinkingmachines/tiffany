package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thinkingmachines/tiffany/pkg/services"
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
		size, _ := cmd.Flags().GetIntSlice("size")
		if len(size) != 2 {
			return errors.New("Requires a size in the form {L},{W}")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		coordinate := []string{args[0], args[1]}

		// Get flags
		zoom, _ := cmd.Flags().GetInt("zoom")
		size, _ := cmd.Flags().GetIntSlice("size")
		pngPath, _ := cmd.Flags().GetString("pngPath")
		tiffPath, _ := cmd.Flags().GetString("tiffPath")
		jsonPath, _ := cmd.Flags().GetString("jsonPath")

		services.RunPipeline(coordinate, zoom, size, pngPath, tiffPath, jsonPath)
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceP("size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().String("pngPath", "tiffany.out/png/", "path to save GSM images in PNG")
	rootCmd.PersistentFlags().String("tiffPath", "tiffany.out/tiff/", "path to save GSM images in TIFF")
	rootCmd.PersistentFlags().String("jsonPath", "tiffany.out/json/", "path to save GeoJSON labels")
	rootCmd.MarkFlagRequired("coordinate")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
