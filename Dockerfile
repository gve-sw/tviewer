FROM golang:latest

WORKDIR /go/src/github.com/sfloresk/tviewer
COPY . .

run export GOBIN=/go/bin/ && \
    go get github.com/gorilla/mux && \
    go get github.com/nleiva/xrgrpc && \
    go install /go/src/github.com/sfloresk/tviewer/main.go


EXPOSE 9090

WORKDIR /go/

CMD ["/go/bin/main"]
