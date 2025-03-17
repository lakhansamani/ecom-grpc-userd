# Build Stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary with static linking (Alpine compatible)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o userd ./main.go

# Final Runtime Stage (Alpine)
FROM alpine:latest

WORKDIR /app

# Install certificates (required for HTTPS calls)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/userd .

# Expose gRPC port
EXPOSE 50051

# Expose metrics port
EXPOSE 9091

# Run the application
CMD ["./userd"]
