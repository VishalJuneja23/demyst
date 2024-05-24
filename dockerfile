
FROM golang:1.22.3 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-binary

# Production phase
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /build/go-binary .
ENTRYPOINT [ "/app/go-binary"]