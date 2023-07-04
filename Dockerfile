FROM golang:1.20.3-alpine3.17 as builder

COPY . /build/
WORKDIR /build/

RUN go build -o ports-service ./cmd/service
RUN go build -o ports-ingestor ./cmd/ingestor

FROM alpine:3.17 as service
WORKDIR /app
COPY --from=builder /build/ports-service ./

FROM alpine:3.17 as ingestor
WORKDIR /app
COPY --from=builder /build/ports-ingestor ./
