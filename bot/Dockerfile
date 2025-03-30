FROM golang:1.20-alpine AS builder

ENV CGO_ENABLED=1
RUN apk update && apk add --no-cache build-base openssl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o vote-bot cli/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/vote-bot /app/vote-bot
RUN chmod +x /app/vote-bot
EXPOSE 8081
ENTRYPOINT ["/app/vote-bot"]
