# ipsec_exporter
Exporter for strongSwan via vici

This repository is to monitor a running strongSwan via the [VICI](https://wiki.strongswan.org/projects/strongswan/wiki/Vici).

It connects to the socket and queries the "list-sas" command, in response gets a list of streamed "list-sa" events from the strongswan core back.
These streamed events are then marshalled to a go slice, which is then analyzed for the metrics itself.

## Usage

### Customizing

By default, there is no customizing.

### Mounts

#### Socket

The programm needs to read the socket at */var/run/charon.vici*. 
This is the default path for strongswan to place the unix socket, so if you did not change it, that should work.
However, due to the containerization you need to mount that file explicitly, or mount */var/run* of the strongSwan to the container.

You will need read and write access to this socket, as the program is issuing a command and reading the answer.

## Error Handling

The program simply logs any errors. If you are unable to connect to the vici socket, this will produce a log entry at every scrape, not more or less.
The scrape itself will just return a lot of empty metrics, but be up, as the prometheus part is fully functional.
In normal mode this should never happen.

## Prometheus

The official prometheus port is 9814. 
It is exposed via Dockerfile and is in the serve command of the webserver.
The port is also published at the [github list of exporters](https://github.com/prometheus/prometheus/wiki/Default-port-allocations)

## Grafana

This are some panels i use to monitor my ongoing connections.

### traffic monitor
Query A (Legend: inbound {{child_name}}):

sum(increase(strongswan_sa_bytes_inbound{ike_name="${ike}"}[1m])) by (child_name)

Query B (Legend: outbound {{child_name}}):

sum(increase(strongswan_sa_bytes_outbound{ike_name="${ike}"}[1m])) by (child_name)

<img width="787" alt="image" src="https://user-images.githubusercontent.com/5329497/113253664-d84bf000-92c5-11eb-9635-b9fb182150f9.png">



### time until rekey

Query A (Legend: Child ID {{child_name}} {{child_id}})

strongswan_sa_rekey_second{ike_name="${ike}"}

<img width="787" alt="image" src="https://user-images.githubusercontent.com/5329497/113253628-c66a4d00-92c5-11eb-9a1b-3e3d8bd92641.png">


## WIP

This repository should be considered as Work in Progress, as there may be changes to the number of metrics.


This go-programm connects to a socket (unix socket or tcp socket) representing a vici connection.
With this connection the strongswan system is queried at every scrape for different statistics.

This informations can be consumed by prometheus and displayed by any tool behind prometheus.
