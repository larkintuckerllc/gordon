package gordon

import (
	"log"
	"net/http"
	"os"
)

var projectId string
var zoneName string
var dnsName string

func Execute(p string, z string, d string) error {
	projectId = p
	zoneName = z
	dnsName = d
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":"+port, nil)
	return err
}
