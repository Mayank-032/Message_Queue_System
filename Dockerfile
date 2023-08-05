FROM golang:1.16.4-buster AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY main.go .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"

# Install MySQL client library
RUN apt-get update && apt-get install -y default-mysql-client

# Install RabbitMQ client library
# (Replace 'rabbitmq-client-library' with the actual package name if available in the Debian repositories)
RUN apt-get update && apt-get install -y rabbitmq-client-library

CMD ["main"]