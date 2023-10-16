FROM golang:1.19 AS builder

WORKDIR /app

ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct

RUN go env -w GOPROXY=$GOPROXY

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/config/config.yaml ./config/

EXPOSE 8080

CMD ["./app"]
