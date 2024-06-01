FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o spy ./cmd/server/main.go

FROM golang:1.18

WORKDIR /app

COPY --from=builder /app/spy .

CMD ["./spy"]
