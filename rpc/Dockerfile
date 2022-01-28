FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /build/zero/rpc/mag rpc/mag.go

CMD ["/build/zero/rpc/mag", "-f", "/build/zero/rpc/etc/mag.yaml"]
