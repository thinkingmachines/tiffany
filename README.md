# tiffany 

A command-line tool for rendering to TIFF any image from Google Static Maps

## Installation

### Compiling from source

You need [go1.11](https://golang.org/doc/go1.11) to be able to compile this from
source. First, clone the repository and enter it:

```s
$ git clone git@github.com:thinkingmachines/tiffany.git
$ cd tiffany
```

Then get the dependencies and build the project:

```s
$ go get
$ go build .
```

Optionally, you can also install `tiffany` inside your system

```s
$ go install
```

### Getting the binaries

Alternatively, you can simply get the latest binaries from our
[Releases](https://github.com/thinkingmachines/tiffany/releases) tab. Make sure
to download the one compatible to your system.


## Usage

### Authentication

Tiffany requires a [Google Static Maps API
Key](https://developers.google.com/maps/documentation/maps-static/intro#get-a-key). Generate one and store it inside an `.env` file in your project directory:

```s
# .env
API_KEY="<your API key here>"
```

### Getting images

![Demo](assets/tiffany-demo.gif)

To get images, simply call `tiffany`, and pass it your latitude and longitude:

```s
# tiffany [LATITUDE] [LONGITUDE]
$ tiffany 14.546943935986324 121.01974525389744
```

This will generate a directory, `tiffany.out` where a `*.png` and its
corresponding `*.tiff` file is located.

In case you don't want georeferenced images and prefer plain-old PNG images,
then simply pass the `--without-reference` flag:

```s
$ tiffany 14.546943935986324 121.01974525389744 --without-reference
```

You can find more options by running `tiffany --help`

## Contributing

Simply fork this repository and [make a Pull
Request](https://help.github.com/en/articles/creating-a-pull-request)! I'm
open to any kind of contribution, but I'd definitely appreciate:

- Implementation of new features 
- Writing documentation
- Testing

Also, we have a
[CONTRIBUTING.md](https://github.com/thinkingmachines/tiffany/blob/master/CONTRIBUTING.md)
and a [Code of
Conduct](https://github.com/thinkingmachines/tiffany/blob/master/CODE_OF_CONDUCT.md),
so please check that one out!

## License

MIT License (c) 2019,  Thinking Machines Data Science
