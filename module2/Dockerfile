FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

COPY httpServer httpServer

EXPOSE 8088
CMD ["httpServer"]

