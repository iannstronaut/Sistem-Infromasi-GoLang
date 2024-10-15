FROM golang:1.23.2-alpine

WORKDIR /redis_docker

COPY . .

RUN go get github.com/go-redis/redis/v8

RUN go mod tidy

RUN go build -o main .
CMD ["./main"]