FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o main main.go

FROM scratch

COPY --from=builder /app/main /main

CMD ["/main"]
