package geonames

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/palantir/stacktrace"
)

type Feature struct {
	GeonameId        int       // geonameid         : integer id of record in geonames database
	Name             string    // name              : name of geographical point (utf8) varchar(200)
	AsciiName        string    // asciiname         : name of geographical point in plain ascii characters, varchar(200)
	AlternateNames   []string  // alternatenames    : alternatenames, comma separated, ascii names automatically transliterated, convenience attribute from alternatename table, varchar(10000)
	Latitude         float64   // latitude          : latitude in decimal degrees (wgs84)
	Longitude        float64   // longitude         : longitude in decimal degrees (wgs84)
	Class            string    // feature class     : see http://www.geonames.org/export/codes.html, char(1)
	Code             string    // feature code      : see http://www.geonames.org/export/codes.html, varchar(10)
	CountryCode      string    // country code      : ISO-3166 2-letter country code, 2 characters
	Cc2              string    // cc2               : alternate country codes, comma separated, ISO-3166 2-letter country code, 200 characters
	Admin1Code       string    // admin1 code       : fipscode (subject to change to iso code), see exceptions below, see file admin1Codes.txt for display names of this code; varchar(20)
	Admin2Code       string    // admin2 code       : code for the second administrative division, a county in the US, see file admin2Codes.txt; varchar(80)
	Admin3Code       string    // admin3 code       : code for third level administrative division, varchar(20)
	Admin4Code       string    // admin4 code       : code for fourth level administrative division, varchar(20)
	Population       int       // population        : bigint (8 byte int)
	Elevation        int       // elevation         : in meters, integer
	Dem              int       // dem               : digital elevation model, srtm3 or gtopo30, average elevation of 3''x3'' (ca 90mx90m) or 30''x30'' (ca 900mx900m) area in meters, integer. srtm processed by cgiar/ciat.
	TimeZone         string    // timezone          : the timezone id (see file timeZone.txt) varchar(40)
	ModificationDate time.Time // modification date : date of last modification in yyyy-MM-dd format
}

// FetchCountryFeatures return list of features for a country
func FetchCountryFeatures(countryIso2Code string, useCache bool) ([]*Feature, error) {
	geonamesUrl := fmt.Sprintf(geonamesUrls["features"], countryIso2Code)
	geonamesZipFile, err := downloadFile(geonamesUrl, useCache)
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file (%s) with features", geonamesUrl)
	}

	// handles 404 error
	if geonamesZipFile == "" {
		return []*Feature{}, nil
	}

	geonamesDir, err := unZip(geonamesZipFile)
	if err != nil {
		return nil, stacktrace.Propagate(err, "unzip geonames archive file")
	}

	return parseFeatures(path.Join(geonamesDir, countryIso2Code+".txt"))
}

func parseFeatures(filename string) ([]*Feature, error) {
	features := make([]*Feature, 0)

	err := parseCsvFile(filename, 0, '\t', '#', func(raw []string) bool {
		if len(raw) != 19 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return false
		}

		geonameId, _ := strconv.Atoi(raw[0])
		latitude, _ := strconv.ParseFloat(raw[4], 64)
		longitude, _ := strconv.ParseFloat(raw[5], 64)
		population, _ := strconv.Atoi(raw[14])
		elevation, _ := strconv.Atoi(raw[15])
		dem, _ := strconv.Atoi(raw[16])
		modificationDate, _ := time.Parse("2006-02-01", raw[18])

		features = append(features, &Feature{
			GeonameId:        geonameId,
			Name:             raw[1],
			AsciiName:        raw[2],
			AlternateNames:   strings.Split(raw[3], ","),
			Latitude:         latitude,
			Longitude:        longitude,
			Class:            raw[6],
			Code:             raw[7],
			CountryCode:      raw[8],
			Cc2:              raw[9],
			Admin1Code:       raw[10],
			Admin2Code:       raw[11],
			Admin3Code:       raw[12],
			Admin4Code:       raw[13],
			Population:       population,
			Elevation:        elevation,
			Dem:              dem,
			TimeZone:         raw[17],
			ModificationDate: modificationDate,
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse csv file with features")
	}

	return features, nil
}
