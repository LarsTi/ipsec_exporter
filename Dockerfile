FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.mod
RUN go mod tidy
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main /app/main
# Reserved Exporter Port for this exporter
EXPOSE 9814 
CMD ["/app/main"]

