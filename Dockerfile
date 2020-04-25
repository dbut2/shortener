FROM golang:1.8-alpine
RUN go build cmd/app/shortener.go
EXPOSE 8080
CMD ["./shortener"]