FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GIN_MODE=release
RUN go build -o secretgen

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/secretgen .

EXPOSE 8080

CMD ["/app/secretgen"]