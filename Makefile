build: gdal
	go get github.com/mitchellh/gox
	OUTPUT=dist/{{.Dir}}\_${DRONE_TAG}\_{{.OS}}_{{.Arch}}
	gox -os="linux" -arch="amd64" -output=${OUTPUT}
	ls dist/

gdal:
	ci/install_gdal.sh
