FROM golang:1.21.6-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o service .

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/service /service
ENTRYPOINT ["./service"]