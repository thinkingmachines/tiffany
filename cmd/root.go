package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thinkingmachines/tiffany/pkg/auth"
	"github.com/thinkingmachines/tiffany/pkg/services"
)

var rootCmd = &cobra.Command{
	Use:   "tiffany",
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
		coordinates, _ := cmd.Flags().GetStringSlice("coordinate")
		if len(coordinates) != 2 {
			fmt.Println(len(coordinates))
			return errors.New("Requires a coordinate in the form {lat},{lon}")
		}
		size, _ := cmd.Flags().GetIntSlice("size")
		if len(size) != 2 {
			return errors.New("Requires a size in the form {L},{W}")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		coordinate, _ := cmd.Flags().GetStringSlice("coordinate")
		zoom, _ := cmd.Flags().GetInt("zoom")
		size, _ := cmd.Flags().GetIntSlice("size")
		fmt.Println(coordinate, zoom, size)
	},
}

func init() {
	rootCmd.Flags().StringSliceP("coordinate", "c", []string{"", ""}, "center coordinate {lat},{lon} (required)")
	rootCmd.PersistentFlags().IntP("zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceP("size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().String("pngPath", "./tiffany.out/png/", "path to save GSM images in PNG")
	rootCmd.PersistentFlags().String("tiffPath", "./tiffany.out/tiff/", "path to save GSM images in TIFF")
	rootCmd.PersistentFlags().String("jsonPath", "./tiffany.out/json/", "path to save GeoJSON labels")
	rootCmd.MarkFlagRequired("coordinate")
}

func pipeline(coordinate []string, zoom int, size []int, pngPath string, tiffPath string, jsonPath string) {
	client := auth.GetStaticMapsClient()
	gsmImage := services.GetGSMImage(client, coordinate, zoom, size)
	services.SaveImagePNG(gsmImage, pngPath)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
