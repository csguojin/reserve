FROM golang:1.19

WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]