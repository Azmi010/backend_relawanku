# phase 1
FROM golang:1.22.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# phase 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
COPY serviceAccountKey.json serviceAccountKey.json
RUN apk update && apk add --no-cache ca-certificates
EXPOSE 8000
CMD ["./main"]