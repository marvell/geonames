package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchAlternateNames(t *testing.T) {
	alternateNames, err := FetchAlternateNames(true)
	if err != nil {
		t.Fatal(err)
	}

	firstRecord := alternateNames[0]
	assert.Equal(t, 3556002, firstRecord.Id, "wrong ID for first record")
}
