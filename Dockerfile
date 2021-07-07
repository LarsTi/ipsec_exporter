FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod go.mod
RUN go get github.com/prometheus/client_golang/prometheus && \
	go get github.com/prometheus/client_golang/prometheus/promhttp && \
	go get github.com/strongswan/govici/vici
#	go mod download
COPY app /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main /app/main
# Reserved Exporter Port for this exporter
EXPOSE 9814 
CMD ["/app/main"]

