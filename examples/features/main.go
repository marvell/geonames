package main

import (
	"log"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	usFeatures, err := geonames.FetchCountryFeatures("US", true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("count: %d", len(usFeatures))
	if len(usFeatures) > 0 {
		log.Printf("%#v", usFeatures[:5])
	}
}
