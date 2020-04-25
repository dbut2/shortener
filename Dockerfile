FROM golang:latest
COPY . .
RUN go build cmd/app/shortener.go
EXPOSE 8080
CMD ["./shortener"]