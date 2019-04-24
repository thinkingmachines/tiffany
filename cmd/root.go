// Copyright 2019 Thinking Machines Data Science. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root
// for license information.

// Package cmd contains all the helper functions, handlers, and command-line methods
// for building the tiffany command-line interface.
package cmd

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/thinkingmachines/tiffany/pkg/pipeline"
)

var rootCmd = &cobra.Command{
	Use:   "tiffany LATITUDE LONGITUDE",
	Short: "tiffany is a tool for rendering to TIFF any image from Google Static Maps",
	Long: `

	     _   _  __  __               
	    | |_(_)/ _|/ _|__ _ _ _ _  _ 
	    |  _| |  _|  _/ _  | ' \ || |
	     \__|_|_| |_| \__,_|_||_\_, |
				    |__/ 

    Render to TIFF any Google Static Maps (GSM) image
       (c) Thinking Machines Data Science, 2019
		Version: v1.0.0-alpha.2`,
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
	Version: "v1.0.0-alpha.2",
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		coordinate := []string{args[0], args[1]}
		initLogger(verbosity)
		client := pipeline.GetStaticMapsClient(env)
		skip := pipeline.Run(client, coordinate, zoom, size, path, noRef, wtLbl, force)
		log.WithFields(log.Fields{
			"lat":     coordinate[0],
			"lon":     coordinate[1],
			"skipped": skip,
		}).Info("Single job done!")
	},
}

var zoom int
var size []int
var path string
var wtLbl string
var noRef bool
var skipFirst bool
var force bool
var verbosity int
var env string

func init() {
	// Add sub-commands
	rootCmd.AddCommand(batchCmd)

	// Define flags for ROOT command
	rootCmd.PersistentFlags().IntVarP(&zoom, "zoom", "z", int(16), "zoom level")
	rootCmd.PersistentFlags().IntSliceVarP(&size, "size", "s", []int{400, 400}, "image size in pixels {L},{W}")
	rootCmd.PersistentFlags().StringVar(&path, "path", "tiffany.out/", "path to save output artifacts")
	rootCmd.PersistentFlags().StringVar(&wtLbl, "with-labels", "", "path to the label's ESRI Shapefile")
	rootCmd.PersistentFlags().BoolVar(&noRef, "without-reference", false, "do not georeference")
	rootCmd.PersistentFlags().BoolVar(&force, "force", false, "download satellite image even if it exists")
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbosity", "v", "set verbosity")
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", ".tiffany.env", "path to .tiffany.env file")

	// Define flags for `batch` command
	batchCmd.Flags().BoolVar(&skipFirst, "skip-header-row", false, "skip header row")
}

func initLogger(verbosity int) {
	if verbosity == 1 {
		log.SetLevel(log.DebugLevel)
	} else if verbosity > 1 {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
