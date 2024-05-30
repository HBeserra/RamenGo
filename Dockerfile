FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app

FROM scratch

COPY --from=builder /app/app /app

EXPOSE 3000

ENTRYPOINT ["/app"]