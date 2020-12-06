build: gdal
	go get -v github.com/mitchellh/gox
	export PATH=${PATH}:${GOPATH}/bin
	${GOPATH}/bin/gox -os="linux darwin windows" 		\
	    -arch="amd64" 					\
	    -output="dist/{{.Dir}}_$(TAG)_{{.OS}}_{{.Arch}}"    \
	    -cgo
	ls dist/

gdal:
	ci/install_gdal.sh
