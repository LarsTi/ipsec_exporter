# ipsec_exporter
Exporter for strongSwan via vici

This repository is to monitor a running strongSwan via the (VICI) [https://wiki.strongswan.org/projects/strongswan/wiki/Vici].

It connects to the socket and queries the "list-sas" command, in response gets a list of streamed "list-sa" events from the strongswan core back.
These streamed events are then marshalled to a go slice, which is then analyzed for the metrics itself.

## Customizing

By default, there is no customizing.

## Prometheus

The official prometheus port is 9814. 
It is exposed via Dockerfile and is in the serve command of the webserver.
The port is also published at the (github list of exporters) [https://github.com/prometheus/prometheus/wiki/Default-port-allocations]

## Grafana

This are some panels i use to monitor my ongoing connections.

### traffic monitor
Query A (Legend: inbound {{child_name}}):

sum(increase(strongswan_sa_bytes_inbound{ike_name="${ike}"}[1m])) by (child_name)

Query B (Legend: outbound {{child_name}}):

sum(increase(strongswan_sa_bytes_outbound{ike_name="${ike}"}[1m])) by (child_name)

### time until rekey

Query A (Legend: Child ID {{child_name}} {{child_id}})

strongswan_sa_rekey_second{ike_name="${ike}"}

## WIP

This repository should be considered as Work in Progress, as there may be changes to the number of metrics.


This go-programm connects to a socket (unix socket or tcp socket) representing a vici connection.
With this connection the strongswan system is queried at every scrape for different statistics.

This informations can be consumed by prometheus and displayed by any tool behind prometheus.
