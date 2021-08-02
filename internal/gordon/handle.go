package gordon

import (
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
	method, instanceProjectId, zone, instanceName, err := parse(&body)
	if err != nil {
		log.Printf("parse: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	switch *method {
	case Insert, Start:
		// Insert: Cloud Pub/Sub set with 30 second "Minimum backoff duration" minimizes errors
		ip, err := getIP(*instanceProjectId, *zone, *instanceName)
		if err != nil {
			log.Printf("getIp: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		err = createRecord(projectId, *instanceName, *ip)
		if err != nil {
			log.Printf("createRecords: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	case Stop:
		err = deleteRecord(projectId, *instanceName)
		if err != nil {
			log.Printf("deleteRecords: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	case Delete:
		// Delete: Guard in case Instance was stopped first
		err = getRecord(projectId, *instanceName)
		if err != nil {
			return
		}
		err = deleteRecord(projectId, *instanceName)
		if err != nil {
			log.Printf("deleteRecords: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}
}
