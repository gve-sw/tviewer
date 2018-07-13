FROM golang:latest

WORKDIR /go
COPY . .

run export GOBIN=/go/bin/ && \
    go get github.com/cisco-gve/tviewer && \
    go install github.com/cisco-gve/tviewer

EXPOSE 9090

WORKDIR /go/

CMD ["/go/bin/tviewer"]
