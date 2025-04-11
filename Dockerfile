FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o book-service cmd/book-service/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/book-service .
EXPOSE 8080
CMD ["./book-service"]