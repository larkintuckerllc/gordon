package gordon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	// TODO: AUTHENTICATE ORIGIN OF PUB/SUB
	// https://cloud.google.com/run/docs/tutorials/pubsub#integrating-pubsub
	var m PubSubMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// TODO: HANDLE
	data := string(m.Message.Data)
	log.Println(data)
}
