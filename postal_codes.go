package geonames

import (
	"fmt"
	"path"
	"strconv"

	"github.com/palantir/stacktrace"
)

// PostalCode struct describe structure of Geonames Postal Code files
type PostalCode struct {
	CountryIso2Code string  // country code      : iso country code, 2 characters
	PostalCode      string  // postal code       : varchar(20)
	PlaceName       string  // place name        : varchar(180)
	AdminName1      string  // admin name1       : 1. order subdivision (state) varchar(100)
	AdminCode1      string  // admin code1       : 1. order subdivision (state) varchar(20)
	AdminName2      string  // admin name2       : 2. order subdivision (county/province) varchar(100)
	AdminCode2      string  // admin code2       : 2. order subdivision (county/province) varchar(20)
	AdminName3      string  // admin name3       : 3. order subdivision (community) varchar(100)
	AdminCode3      string  // admin code3       : 3. order subdivision (community) varchar(20)
	Latitude        float64 // latitude          : estimated latitude (wgs84)
	Longitude       float64 // longitude         : estimated longitude (wgs84)
	Accuracy        int     // accuracy          : accuracy of lat/lng from 1=estimated to 6=centroi}
}

// FetchPostalCodes returns list of languages
func FetchCountryPostalCodes(countryIso2Code string, useCache bool) ([]PostalCode, error) {
	geonamesZipFile, err := downloadFile(fmt.Sprintf(geonamesUrls["postal_codes"], countryIso2Code), useCache)
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file with postal codes")
	}

	geonamesDir, err := unZip(geonamesZipFile)
	if err != nil {
		return nil, stacktrace.Propagate(err, "unzip geonames archive file")
	}

	filename := path.Join(geonamesDir, countryIso2Code+".txt")

	postalCodes := make([]PostalCode, 0)
	err = parseCsvFile(filename, 0, '\t', '#', func(raw []string) bool {
		if len(raw) != 12 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return false
		}

		latitude, _ := strconv.ParseFloat(raw[9], 64)
		longitude, _ := strconv.ParseFloat(raw[10], 64)
		accuracy, _ := strconv.Atoi(raw[11])

		postalCodes = append(postalCodes, PostalCode{
			CountryIso2Code: raw[0],
			PostalCode:      raw[1],
			PlaceName:       raw[2],
			AdminName1:      raw[3],
			AdminCode1:      raw[4],
			AdminName2:      raw[5],
			AdminCode2:      raw[6],
			AdminName3:      raw[7],
			AdminCode3:      raw[8],
			Latitude:        latitude,
			Longitude:       longitude,
			Accuracy:        accuracy,
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse csv file with postal codes")
	}

	return postalCodes, nil
}
