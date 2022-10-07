ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder
COPY . /$GOPATH/src/gin-api/
WORKDIR /$GOPATH/src/gin-api/
#setup environment variable and proxy
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://proxy.golang.org,direct
#generate go.mod
RUN go mod init
#add necessary dependencies and remove dependencies
RUN go mod tidy
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/weppos/dnscaa
RUN go get -u golang.org/x/net/webdav
RUN go get -u github.com/shuque/dane
RUN go build index.go
EXPOSE 8080:8080
CMD ["./index"]

