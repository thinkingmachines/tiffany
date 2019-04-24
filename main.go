// Copyright 2019 Thinking Machines Data Science. All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root
// for license information.

/*
Tiffany is command-line tool for rendering to TIFF any image from Google Static Maps.

It downloads, georeferences, and labels any satellite image from the Static
Maps API. You can use this to prepare labeled data for downstream tasks such as
in computer vision (object detection, semantic segmentation, etc.)

Installation

You can get the binaries from our Github releases:
https://github.com/thinkingmachines/tiffany/releases

Or, you can compile this from source by cloning the repository and building it:

    $ git clone git@github.com:thinkingmachines/tiffany.git
    $ cd tiffany
    $ go get
    $ go build .


Usage

Usage instructions can be found in the README:
https://github.com/thinkingmachines/tiffany/blob/master/README.md

Contributing

Simply fork the Github repository and make a Pull Request. We're open to any
kind of contribution, but we'd definitely appreciate (1) implementation of new
features (2) writing documentation and (3) testing.

License

MIT License (c) 2019, Thinking Machines Data Science
*/
package main

import (
	"github.com/thinkingmachines/tiffany/cmd"
)

func main() {
	cmd.Execute()
}
