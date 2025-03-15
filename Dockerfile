# Stage 1: Build the Go binary
FROM golang:1.24.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY main.go ./

# Build the Go app (statically linked binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o web-service

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/web-service .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./web-service"]
