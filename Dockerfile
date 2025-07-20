FROM golang:1.23.6-bullseye AS builder

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o graphql-api .

# build
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/graphql-api .

EXPOSE 8080

CMD ["./graphql-api"]
