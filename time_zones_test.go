package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchTimeZones(t *testing.T) {
	timeZones, err := FetchTimeZones()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, timeZones, "empty time zones slice")

	abidjan := timeZones[0]
	assert.Equal(t, "CI", abidjan.CountryIso2Code, "wrong country for Abidjan")
	assert.Equal(t, "Africa/Abidjan", abidjan.TimeZoneId, "wrong id for Abidjan")
}
