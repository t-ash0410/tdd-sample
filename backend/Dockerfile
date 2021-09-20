FROM golang:latest

RUN apt-get update && \
    apt-get install -y curl git procps

RUN go get -v golang.org/x/tools/gopls && \
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs && \
    go get -v github.com/ramya-rao-a/go-outline && \
    go get -v github.com/fatih/gomodifytags && \
    go get -v github.com/josharian/impl && \
    go get -v github.com/haya14busa/goplay/cmd/goplay && \
    go get -v github.com/go-delve/delve/cmd/dlv