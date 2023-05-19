FROM golang:1.19 AS builder

WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct

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
