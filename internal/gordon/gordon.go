package gordon

import (
	"log"
	"net/http"
	"os"
)

func Execute() error {
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
