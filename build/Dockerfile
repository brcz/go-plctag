FROM ubuntu:20.04 as libplc-builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    apt-utils \
    build-essential \
    cmake \
    ubuntu-dev-tools \
    wget \
    mlocate \
    autoconf autogen automake \
    pkg-config libtool libcurl3-openssl-dev libssl-dev \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir /artifact

COPY . /go/src/github.com/brcz/go-plctag

WORKDIR /go/src/github.com/brcz/go-plctag

CMD ["bash", "-c","build/build-lib.sh"]
