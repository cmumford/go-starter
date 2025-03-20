# Stage 1: Build the Go binary
FROM golang:1.24.1 AS builder

ARG GIT_SHA1=prod
ENV GIT_COMMIT_ID=$GIT_SHA1

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app (statically linked binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o web-service

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /usr/src/app/web-service .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./web-service"]
