package main

import (
	"fmt"

	"github.com/marvell/geonames"
)

func main() {
	geonames.EnableDebugMode()

	countries, err := geonames.FetchCountries(true)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	alternateNames, err := geonames.FetchAlternateNames(true)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	country := countries[0]
	fmt.Printf("Country: %s\nAlternate name: ", country.Name)

	for i := range alternateNames {
		if alternateNames[i].GeonameId == country.GeonameId {
			fmt.Printf("%s, ", alternateNames[i].Name)
		}
	}
}
