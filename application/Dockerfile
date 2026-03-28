FROM golang:1.25.5 AS builder
#
# Build stage
#
WORKDIR /usr/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o webapp ./cmd/web/*.go

FROM alpine:latest
#
# Package stage
#
WORKDIR /usr/app

COPY --from=builder /usr/app/webapp .

EXPOSE 4000

CMD ["./webapp"]
