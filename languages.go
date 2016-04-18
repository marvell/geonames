package geonames

import (
	"github.com/palantir/stacktrace"
)

// Language struct describe structure of iso-languagecodes.txt file entry
type Language struct {
	Iso639_3     string // ISO 639-3
	Iso639_2     string // ISO 639-2
	Iso639_1     string // ISO 639-1
	LanguageName string // Language Name
}

// FetchLanguages returns list of languages
func FetchLanguages(useCache bool) ([]*Language, error) {
	geonamesFile, err := downloadFile(geonamesUrls["languages"], useCache)
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file with languages")
	}

	languages := make([]*Language, 0)

	err = parseCsvFile(geonamesFile, 1, '\t', '#', func(raw []string) bool {
		if len(raw) != 4 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return true
		}

		languages = append(languages, &Language{
			Iso639_3:     raw[0],
			Iso639_2:     raw[1],
			Iso639_1:     raw[2],
			LanguageName: raw[3],
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse cav file with languages")
	}

	return languages, nil
}
