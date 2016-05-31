package geonames

import (
	"strconv"

	"github.com/palantir/stacktrace"
)

// Country struct describe structure of countryInfo.txt file entry
type Country struct {
	Iso2Code           string  // ISO
	Iso3Code           string  // ISO3
	IsoNumeric         string  // ISO-Numeric
	Fips               string  // fips
	Name               string  // Country
	Capital            string  // Capital
	Area               float64 // Area(in sq km)
	Population         uint    // Population
	Continent          string  // Continent
	Tld                string  // tld
	CurrencyCode       string  // CurrencyCode
	CurrencyName       string  // CurrencyName
	Phone              string  // Phone
	PostalCodeFormat   string  // Postal Code Format
	PostalCodeRegex    string  // Postal Code Regex
	Languages          string  // Languages
	GeonameId          int     // geonameid
	Neighbours         string  // neighbours
	EquivalentFipsCode string  // EquivalentFipsCode
}

// FetchCountries returns list of countries
func FetchCountries(useCache bool) ([]*Country, error) {
	geonamesFile, err := downloadFile(geonamesUrls["countries"], useCache)
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file with countries")
	}

	countries := make([]*Country, 0)

	err = parseCsvFile(geonamesFile, 0, '\t', '#', func(raw []string) bool {
		if len(raw) != 19 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return false
		}

		area, _ := strconv.ParseFloat(raw[6], 64)
		population, _ := strconv.ParseUint(raw[7], 10, 32)
		geonamesId, _ := strconv.Atoi(raw[16])

		countries = append(countries, &Country{
			Iso2Code:           raw[0],
			Iso3Code:           raw[1],
			IsoNumeric:         raw[2],
			Fips:               raw[3],
			Name:               raw[4],
			Capital:            raw[5],
			Area:               area,
			Population:         uint(population),
			Continent:          raw[8],
			Tld:                raw[9],
			CurrencyCode:       raw[10],
			CurrencyName:       raw[11],
			Phone:              raw[12],
			PostalCodeFormat:   raw[13],
			PostalCodeRegex:    raw[14],
			Languages:          raw[15],
			GeonameId:          geonamesId,
			Neighbours:         raw[17],
			EquivalentFipsCode: raw[18],
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse cav file with countries")
	}

	return countries, nil
}
