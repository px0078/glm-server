# Choose whatever you want, version >= 1.16
FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . ./
RUN go mod download


CMD ["air", "-c", "/app/.air.toml"]