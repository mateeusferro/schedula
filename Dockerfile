FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/ cmd/main.go


FROM alpine AS runner

WORKDIR /app

COPY --from=builder /app/bin/main ./main
COPY .env .env
CMD ["/app/main"]