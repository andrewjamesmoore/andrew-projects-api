FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o graphql-api ./server.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/graphql-api .

EXPOSE 8080
CMD ["./graphql-api"]
