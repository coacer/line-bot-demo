FROM golang:1.17.0-alpine3.14

COPY ./src .
RUN go build -o main main.go
ENTRYPOINT ./main
