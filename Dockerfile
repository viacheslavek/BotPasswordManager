FROM golang:alpine

WORKDIR /go/src/myapp

COPY . .

RUN go mod download

RUN go build -o myapp cmd/main.go

EXPOSE 8080

CMD ["./myapp", "-tg_token", "${TELEGRAM_TOKEN}"]
