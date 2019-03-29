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
which gdal-config || exit ${LINENO}
cd /usr/lib/pkgconfig || exit ${LINENO}
[ -r gdal.pc ] && exit ${LINENO}
curl -ks 'https://gist.githubusercontent.com/nicerobot/5160658/raw/gdal.pc.template' \
| sed -e 's/^/echo /' \
| bash - \
> ${1:-gdal.pc}
cat /usr/lib/pkgconfig/gdal.pc
