FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod go.mod
RUN go mod download
COPY app /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main /app/main
CMD ["/app/main"]
