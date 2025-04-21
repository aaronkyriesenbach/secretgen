FROM golang:1.24-alpine AS base

WORKDIR /app

COPY . .

RUN go mod download

ENV GIN_MODE=release
RUN go build -o secretgen

EXPOSE 8080

CMD ["/app/secretgen"]