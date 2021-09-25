FROM golang:latest

RUN apt-get update && \
    apt-get install -y curl git procps unzip

RUN go get -v golang.org/x/tools/gopls && \
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs && \
    go get -v github.com/ramya-rao-a/go-outline && \
    go get -v github.com/fatih/gomodifytags && \
    go get -v github.com/josharian/impl && \
    go get -v github.com/haya14busa/goplay/cmd/goplay && \
    go get -v github.com/go-delve/delve/cmd/dlv && \
    go get -v google.golang.org/grpc && \
    go get -v github.com/golang/protobuf/protoc-gen-go

RUN mkdir /work
WORKDIR /work
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.9.0/protoc-3.9.0-linux-x86_64.zip && \
    unzip protoc-3.9.0-linux-x86_64.zip -d protoc3 && \
    mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/ && \
    rm -rf protoc-3.9.0-linux-x86_64.zip protoc3