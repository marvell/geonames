package geonames

import (
	"path"
	"strconv"

	"github.com/palantir/stacktrace"
)

type AlternateName struct {
	Id              int    // alternateNameId   : the id of this alternate name, int
	GeonameId       int    // geonameid         : geonameId referring to id in table 'geoname', int
	IsoLanguage     string // isolanguage       : iso 639 language code 2- or 3-characters; 4-characters 'post' for postal codes and 'iata','icao' and faac for airport codes, fr_1793 for French Revolution names,  abbr for abbreviation, link for a website, varchar(7)
	Name            string // alternate name    : alternate name or name variant, varchar(200)
	IsPreferredName bool   // isPreferredName   : '1', if this alternate name is an official/preferred name
	IsShortName     bool   // isShortName       : '1', if this is a short name like 'California' for 'State of California'
	IsColloquial    bool   // isColloquial      : '1', if this alternate name is a colloquial or slang term
	IsHistoric      bool   // isHistoric        : '1', if this alternate name is historic and was used in the past
}

// FetchAlternateNames returns list of alternate names
func FetchAlternateNames(useCache bool) ([]AlternateName, error) {
	geonamesZipFile, err := downloadFile(geonamesUrls["alternate_names"], useCache)
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file with alternate names")
	}

	geonamesDir, err := unZip(geonamesZipFile)
	if err != nil {
		return nil, stacktrace.Propagate(err, "unzip geonames archive file")
	}

	alternateNames := make([]AlternateName, 0)

	err = parseCsvFile(path.Join(geonamesDir, "alternateNames.txt"), 0, '\t', '#', func(raw []string) bool {
		if len(raw) != 8 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return false
		}

		// skip links
		if raw[2] == "link" {
			return true
		}

		id, _ := strconv.Atoi(raw[0])
		geonameId, _ := strconv.Atoi(raw[1])
		boolTrue := "1"

		alternateNames = append(alternateNames, AlternateName{
			Id:              id,
			GeonameId:       geonameId,
			IsoLanguage:     raw[2],
			Name:            raw[3],
			IsPreferredName: raw[4] == boolTrue,
			IsShortName:     raw[5] == boolTrue,
			IsColloquial:    raw[6] == boolTrue,
			IsHistoric:      raw[7] == boolTrue,
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse cav file with alternate names")
	}

	return alternateNames, nil
}
