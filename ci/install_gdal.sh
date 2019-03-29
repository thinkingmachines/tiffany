#!/bin/bash

# Install software-properties-common to add repositories
apt-get update -y && apt install -y software-properties-common

# Install GDAL
add-apt-repository -y ppa:ubuntugis/ppa && apt-get update
apt-get update && apt-get install -y \
    curl \
    gdal-bin \
    libgdal-dev

# Include GDAL to path
export CPLUS_INCLUDE_PATH=/usr/include/gdal
export C_INCLUDE_PATH=/usr/include/gdal

# Add gdal.pc in config
curl -ks 'https://gist.githubusercontent.com/nicerobot/5160658/raw/install-gdalpc.sh' | bash -
cat /usr/lib/pkgconfig/gdal.pc
