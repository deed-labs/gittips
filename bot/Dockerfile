FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download
RUN go build -o bin/bot ./cmd/bot/*

EXPOSE 3000

CMD ./bin/bot

