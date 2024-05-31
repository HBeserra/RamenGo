FROM golang:1.22-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app

EXPOSE 3000

ENTRYPOINT ["/app"] 