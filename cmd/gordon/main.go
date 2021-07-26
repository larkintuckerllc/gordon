package main

import (
	"log"

	"github.com/larkintuckerllc/gordon/internal/gordon"
)

func main() {
	err := gordon.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
