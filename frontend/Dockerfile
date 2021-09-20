FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y tzdata && \
    apt-get install -y git unzip clang xserver-xorg pkg-config libgtk-3-dev curl cmake ninja-build

WORKDIR /opt

RUN git clone https://github.com/flutter/flutter.git

