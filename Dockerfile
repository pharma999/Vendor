# Stage 1: Build
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Disable CGO to create static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Minimal runtime
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
