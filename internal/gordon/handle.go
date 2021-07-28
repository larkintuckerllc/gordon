package gordon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	method, projectId, zone, instanceId, err := parse(&body)
	if err != nil {
		log.Printf("parse: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	ip, err := getIP(*projectId, *zone, *instanceId)
	if err != nil {
		log.Printf("getIp: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Println(method)
	fmt.Println(ip)
}
