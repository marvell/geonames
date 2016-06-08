package main

import (
	"log"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	err := geonames.CleanUp()
	if err != nil {
		log.Fatal(err)
	}

	countries, err := geonames.FetchCountries(true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("count: %d", len(countries))
	if len(countries) > 0 {
		log.Printf("%#v\n\n", countries[0])
	}
}
