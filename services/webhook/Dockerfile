FROM golang:1.17.0 as builder

WORKDIR /go/src
ARG WORK_DIR=./services/trigger

COPY $WORK_DIR/go.mod $WORK_DIR/go.sum ./
RUN go mod download

COPY $WORK_DIR  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
ENV GO111MODULE auto
RUN go build \
    -o /go/bin/main

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main
EXPOSE 5000

ENTRYPOINT ["/app/main"]
