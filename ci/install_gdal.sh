#!/bin/bash

apt-get update -y && apt install -y software-properties-common
add-apt-repository -y ppa:ubuntugis/ppa && apt-get update
# add-apt-repository -y ppa:nextgis/ppa
apt-get update && apt-get install -y \
    gdal-bin \
    libgdal-dev

export CPLUS_INCLUDE_PATH=/usr/include/gdal
export C_INCLUDE_PATH=/usr/include/gdal