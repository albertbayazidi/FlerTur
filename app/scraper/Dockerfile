FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o scraper .

FROM alpine:latest
RUN apk add --no-cache chromium ca-certificates
WORKDIR /app
COPY --from=builder /app/scraper .
ENV ROD_BROWSER_BIN=/usr/bin/chromium-browser
CMD ["./scraper"]
