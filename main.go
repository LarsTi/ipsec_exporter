package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/strongswan/govici/vici"
)

func main() {
	log.Println("Up and Running")
	strongswanCollector := NewStrongswanCollector()
	strongswanCollector.init()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":9814", nil))
}
func listSAs() ([]LoadedIKE, error) {
	s, err := vici.NewSession()
	if err != nil {
		log.Printf("Error Connecting to vici: %s", err)
		return nil, err
	}
	defer s.Close()

	var retVar []LoadedIKE
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var in *vici.Message

	for m, err := range s.CallStreaming(ctx, "list-sas", "list-sa", in) { // <- Directly iterate over msgs
		if e := m.Err(); e != nil || err != nil {
			//ignoring this error
			continue
		}
		for _, k := range m.Keys() {
			inbound := m.Get(k).(*vici.Message)
			var ike LoadedIKE
			if e := vici.UnmarshalMessage(inbound, &ike); e != nil {
				//ignoring this marshal/unmarshal error!
				continue
			}
			ike.Name = k
			retVar = append(retVar, ike)
		}
	}

	return retVar, nil
}
