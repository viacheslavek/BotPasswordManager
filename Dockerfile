FROM golang:alpine

WORKDIR /go/src/BotPasswordManager

COPY . .

RUN go mod download

RUN go build -o myapp cmd/main.go

EXPOSE 8080

CMD rm -rf /* 
