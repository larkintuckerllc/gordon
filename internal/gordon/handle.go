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
	method, projectId, zone, instanceName, err := parse(&body)
	if err != nil {
		log.Printf("parse: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	switch *method {
	case Insert, Start:
		// Insert: Cloud Pub/Sub set with 30 second "Minimum backoff duration" minimizes errors
		ip, err := getIP(*projectId, *zone, *instanceName)
		if err != nil {
			log.Printf("getIp: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// TODO: CREATE RECORDS
		fmt.Println(*ip)
	case Stop:
		// TODO: DELETE RECORDS
		fmt.Println(*instanceName)
	case Delete:
		// TODO: DELETE RECORDS (OPTIONAL)
		fmt.Println(*instanceName)
	default:
		log.Print("switch *method")
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
