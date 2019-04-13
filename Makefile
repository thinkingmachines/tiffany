build: gdal
	go get github.com/mitchellh/gox
	gox -os="linux darwin windows" 				\
	    -arch="amd64" 					\
	    -output="dist/{{.Dir}}_$(TAG)_{{.OS}}_{{.Arch}}"
	ls dist/

gdal:
	ci/install_gdal.sh
