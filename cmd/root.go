package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiffany",
	Short: "tiffany is a tool for obtaining TIFF images from Google Static Maps",
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
			return errors.New("Requires a coordinate in the form {lat},{lon}")
		}
		size, _ := cmd.Flags().GetIntSlice("size")
		if len(size) != 2 {
			return errors.New("Requires a size in the form {L},{W}")
		}
		return fmt.Errorf("Invalid argument")

	},
	Run: func(cmd *cobra.Command, args []string) {
		coordinate, _ := cmd.Flags().GetStringSlice("coordinate")
		zoom, _ := cmd.Flags().GetIntSlice("zoom")
		size, _ := cmd.Flags().GetInt("size")
		fmt.Println(coordinate, zoom, size)
	},
}

func init() {
	rootCmd.Flags().StringSliceP("coordinate", "c", []string{"", ""}, "center coordinate {lat},{lon} (required)")
	rootCmd.PersistentFlags().IntP("zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceP("size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.MarkFlagRequired("coordinate")
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
