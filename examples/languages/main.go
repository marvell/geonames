package main

import (
	"log"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	languages, err := geonames.FetchLanguages(true)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("count: %d", len(languages))
	if len(languages) > 0 {
		log.Printf("%#v\n\n", languages[0])
	}
}
