package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thinkingmachines/tiffany/pkg/auth"
	"github.com/thinkingmachines/tiffany/pkg/services"
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
}

func pipeline() {
	gsmClient = auth.GetStaticMapsClient()
	services.GetGSMImage(gsmClient)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
