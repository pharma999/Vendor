# Stage 1: Build Go binary
FROM golang:1.25.1 AS builder
WORKDIR /app

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate swagger docs
RUN /go/bin/swag init -g main.go

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Runtime image
FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Kolkata /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs   # if swagger docs needed
COPY .env .

CMD ["./main"]
