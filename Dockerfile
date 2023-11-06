# Choose whatever you want, version >= 1.16
FROM golang:1.21-alpine

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "main.go"]