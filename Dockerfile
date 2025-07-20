FROM golang:1.23.6-bullseye AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o graphql-api .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/graphql-api .

EXPOSE 8080

CMD ["./graphql-api"]

