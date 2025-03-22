FROM golang:1.24.1 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-starter

FROM alpine:latest
ARG GIT_SHA1=prod
ENV GIT_COMMIT_ID=$GIT_SHA1
WORKDIR /root/
COPY --from=builder /usr/src/app/go-starter .
EXPOSE 8080
CMD ["./go-starter"]
