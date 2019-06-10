FROM ubuntu:18.04 as libplc-builder

ENV DEBIAN_FRONTEND=noninteractive
ENV PACKAGE github.com/cfnorg/edged/internal/edged/app
RUN apt-get update
RUN apt-get install -y build-essential git cmake

RUN mkdir /artifact

COPY . /go/src/github.com/brcz/go-plctag

WORKDIR /go/src/github.com/brcz/go-plctag

CMD ["bash", "-c","build/build-lib.sh"]
