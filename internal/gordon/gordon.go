package gordon

import (
	"log"
	"net/http"
	"os"
)

var pId string
var zName string
var dName string

// Execute starts the HTTP server.
// It returns any error encountered
func Execute(projectId string, zoneName string, dnsName string) error {
	pId = projectId
	zName = zoneName
	dName = dnsName
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
