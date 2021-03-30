# ipsec_exporter
Exporter for strongSwan via swanctl

This go-programm connects to a socket (unix socket or tcp socket) representing a vici connection.
With this connection the strongswan system is queried at every scrape for different statistics.

This informations can be consumed by prometheus and displayed by any tool behind prometheus.
