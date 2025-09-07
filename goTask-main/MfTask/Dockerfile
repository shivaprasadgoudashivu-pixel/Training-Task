FROM golang:1.24.5 AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Create app directory
WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/grool/OrderFlow.grl ./grool/OrderFlow.grl
EXPOSE 8081

ENTRYPOINT [ "./app" ]

