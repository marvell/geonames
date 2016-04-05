package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCountries(t *testing.T) {
	countries, err := FetchCountries(true)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 252, len(countries), "wrong count of countries")

	andorra := countries[0]
	assert.Equal(t, "AD", andorra.Iso2Code, "wrong ISO2 code for Andorra")
}
