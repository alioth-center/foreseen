FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -o foreseen

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/foreseen .

RUN chmod +x foreseen

EXPOSE 8881

ENTRYPOINT ["/app/foreseen"]