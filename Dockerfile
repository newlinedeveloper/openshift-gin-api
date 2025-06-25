# Dockerfile
FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM registry.access.redhat.com/ubi8/ubi
COPY --from=builder /app/app /app
EXPOSE 8080
CMD ["/app"]
