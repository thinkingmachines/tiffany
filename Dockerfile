# Go 1.11.1 cross compiler with GDAL 2.3.2
FROM karalabe/xgo-latest
MAINTAINER Lester James V. Miranda <lj@thinkingmachin.es>

# Install basic dependencies
RUN apt-get update -y && apt install -y software-properties-common
RUN add-apt-repository -y ppa:ubuntugis/ubuntugis-unstable
RUN apt-get update && apt-get install -y \
    gdal-bin \
    libgdal-dev \
    libgdal20

ENV CPLUS_INCLUDE_PATH /usr/include/gdal
ENV C_INCLUDE_PATH /usr/include/gdal
