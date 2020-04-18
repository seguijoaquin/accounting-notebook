FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o app ./cmd/api

EXPOSE 8080

CMD ["./app"]