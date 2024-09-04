# Dockerfile
FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
