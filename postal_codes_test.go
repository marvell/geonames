package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCountryPostalCodes(t *testing.T) {
	postalCodes, err := FetchCountryPostalCodes("AD", true)

	assert.Nil(t, err)
	assert.NotEmpty(t, postalCodes)

	f := postalCodes[0]
	assert.Equal(t, "AD", f.CountryIso2Code)
	assert.Equal(t, "AD100", f.PostalCode)
	assert.Equal(t, "Canillo", f.PlaceName)
}
