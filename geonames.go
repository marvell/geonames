package geonames

var geonamesUrls = map[string]string{
	"countries":       "http://download.geonames.org/export/dump/countryInfo.txt",
	"time_zones":      "http://download.geonames.org/export/dump/timeZones.txt",
	"languages":       "http://download.geonames.org/export/dump/iso-languagecodes.txt",
	"alternate_names": "http://download.geonames.org/export/dump/alternateNames.zip",

	"features_all": "http://download.geonames.org/export/dump/allCountries.zip",
	"features":     "http://download.geonames.org/export/dump/%s.zip",
}
