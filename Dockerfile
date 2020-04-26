FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/app/shortener.go
EXPOSE 8080
CMD ["./shortener"]