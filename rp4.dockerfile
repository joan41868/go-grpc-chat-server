FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build

EXPOSE 9000
ENTRYPOINT ["/app/grpc-chat"]