package geonames

import (
	"strconv"
	"time"

	"github.com/palantir/stacktrace"
)

// TimeZone struct describe structure of timeZones.txt file entry
type TimeZone struct {
	CountryIso2Code string        // CountryCode
	TimeZoneId      string        // TimeZoneId
	GmtOffset       time.Duration // GMT offset 1. Jan 2016
	DstOffset       time.Duration // DST offset 1. Jul 2016
	RawOffset       time.Duration // rawOffset (independant of DST)
}

// FetchTimeZones returns list of time zones
func FetchTimeZones() ([]TimeZone, error) {
	geonamesFile, err := downloadFile(geonamesUrls["time_zones"])
	if err != nil {
		return nil, stacktrace.Propagate(err, "download geonames file with time zones")
	}

	timeZones := make([]TimeZone, 0)

	err = parseCsvFile(geonamesFile, 1, '\t', '#', func(raw []string) bool {
		if len(raw) != 5 {
			logWarn("invalid (%d) count of fields\n\t=> %v", len(raw), raw)

			return true
		}

		gmtOffset, _ := strconv.ParseFloat(raw[2], 64)
		dstOffset, _ := strconv.ParseFloat(raw[3], 64)
		rawOffset, _ := strconv.ParseFloat(raw[4], 64)

		timeZones = append(timeZones, TimeZone{
			CountryIso2Code: raw[0],
			TimeZoneId:      raw[1],
			GmtOffset:       time.Duration(gmtOffset) * time.Hour,
			DstOffset:       time.Duration(dstOffset) * time.Hour,
			RawOffset:       time.Duration(rawOffset) * time.Hour,
		})

		return true
	})

	if err != nil {
		return nil, stacktrace.Propagate(err, "parse cav file with time zones")
	}

	return timeZones, nil
}
