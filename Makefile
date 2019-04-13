build: gdal
	go get github.com/mitchellh/gox
	gox -os="linux" -arch="amd64" -output="dist/{{.Dir}}\_$(DRONE_TAG)\_{{.OS}}_{{.Arch}}"
	ls dist/

gdal:
	ci/install_gdal.sh
