# Stage 1: build
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o server main.go

# Stage 2: runtime
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server /app/server

ENV SCRIPT_DIR=/app/scripts
ENV PORT=8080

EXPOSE 8080

CMD ["/app/server"]
