# Build stage
FROM golang:1.24.4-alpine3.22 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Final minimal stage
FROM alpine:3.22

WORKDIR /app

ENV APP_ENV=prod

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
