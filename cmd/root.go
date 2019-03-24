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
		path, _ := cmd.Flags().GetString("path")
		noRef, _ := cmd.Flags().GetBool("without-reference")

		RunPipeline(coordinate, zoom, size, path, noRef)
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceP("size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().String("path", "tiffany.out/", "path to save output artifacts")
	rootCmd.PersistentFlags().Bool("without-reference", false, "do not georeference")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
