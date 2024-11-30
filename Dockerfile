# Base image with Go 1.22
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server/main.go

# Expose server port
EXPOSE 8080

# Run the application
CMD ["./main"]
