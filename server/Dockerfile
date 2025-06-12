FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY .. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./server/cmd/main.go

FROM alpine:3.18.6
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8082
ENTRYPOINT ["/app/main"]
