package main

import (
	"log"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	alternateNames, err := geonames.FetchAlternateNames(true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("count: %d", len(alternateNames))
	if len(alternateNames) > 0 {
		log.Printf("%#v\n\n", alternateNames[0])
	}
}
