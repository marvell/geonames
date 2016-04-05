package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCountryFeatures(t *testing.T) {
	features, err := FetchCountryFeatures("AD", true)

	assert.Nil(t, err)
	assert.NotEmpty(t, features)

	f := features[0]
	assert.Equal(t, 2986043, f.GeonameId)
	assert.Equal(t, "Pic de Font Blanca", f.Name)
	assert.Equal(t, 42.64991, f.Latitude)
	assert.Equal(t, 1.53335, f.Longitude)
	assert.Equal(t, "Europe/Andorra", f.TimeZone)
}
