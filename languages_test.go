package geonames

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchLanguages(t *testing.T) {
	languages, err := FetchLanguages()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, languages, "empty languages slice")

	afa := languages[0]
	assert.Equal(t, "afa", afa.Iso639_3, "wrong ISO 639 1 code")
}
