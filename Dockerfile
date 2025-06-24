FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -ldflags "-s -w" -o /app/mimer .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mimer .

EXPOSE 8080

CMD ["/app/mimer"]
