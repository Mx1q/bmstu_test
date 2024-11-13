FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

ENTRYPOINT [ "/app/main" ]