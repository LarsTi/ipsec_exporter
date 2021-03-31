package main

import(
	"log"
	"github.com/strongswan/govici/vici"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Up and Running")
	strongswanCollector := NewStrongswanCollector()
	strongswanCollector.init()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":9814", nil))
}
func listSAs()([]LoadedIKE, error){
	s, err := vici.NewSession()
	if err != nil {
		log.Printf("Error Connecting to vici: %s", err)
		return nil, err
	}

	var retVar []LoadedIKE
	msgs, err := s.StreamedCommandRequest("list-sas", "list-sa", nil)
	if err != nil {
		return retVar, err
	}
	for _,m := range msgs.Messages() {
		if e := m.Err(); e != nil{
			//ignoring this error
			continue
		}
		for _, k := range m.Keys() {
			inbound := m.Get(k).(*vici.Message)
			var ike LoadedIKE
			if e := vici.UnmarshalMessage(inbound, &ike); e != nil {
				//ignoring this marshal/unmarshal errro!
				continue
			}
			ike.Name = k
			retVar = append(retVar, ike)

		}
	}
	return retVar, nil
}
