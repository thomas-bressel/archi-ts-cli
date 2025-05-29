FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /archits
FROM alpine:latest
COPY --from=builder /archits /usr/local/bin/archits
CMD ["archits"]
