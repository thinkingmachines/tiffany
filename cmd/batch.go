// Copyright 2019 Thinking Machines Data Science. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root
// for license information.

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thinkingmachines/tiffany/pkg/pipeline"
)

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
	Version: "v1.0.0-alpha.2",
	Run: func(cmd *cobra.Command, args []string) {
		// Get arguments passed
		csvFile := args[0]
		log.WithFields(log.Fields{
			"file": csvFile,
		}).Info("Batch job successfully started")
		initLogger(verbosity)
		client := pipeline.GetStaticMapsClient(env)
		total, numSkip := pipeline.RunBatch(client, csvFile, skipFirst, zoom, size, path, noRef, wtLbl, force)
		fmt.Println("")
		log.WithFields(log.Fields{
			"total":   total,
			"skipped": numSkip,
		}).Info("Batch job done!")
	},
}
