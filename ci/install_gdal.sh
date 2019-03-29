#!/bin/bash

apt-get update -y && apt install -y software-properties-common
add-apt-repository -y ppa:ubuntugis/ppa && apt-get update
apt-get update && apt-get install -y \
    gdal-bin \
    libgdal-dev

export CPLUS_INCLUDE_PATH=/usr/include/gdal
export C_INCLUDE_PATH=/usr/include/gdal

# Add gdal.pc in config
which gdal-config
cd /usr/lib/pkgconfig
[ -r gdal.pc ]
curl -ks 'https://gist.githubusercontent.com/nicerobot/5160658/raw/gdal.pc.template' \
| sed -e 's/^/echo /' \
| bash - \
> ${1:-gdal.pc}
