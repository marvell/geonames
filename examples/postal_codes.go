package main

import (
	"log"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	postalCodes, err := geonames.FetchCountryPostalCodes("AD", true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("count: %d", len(postalCodes))
	if len(postalCodes) > 0 {
		log.Printf("%#v", postalCodes[:5])
	}
}
