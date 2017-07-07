package geonames

import (
	"strconv"
	"time"
)

const timeZonesURL = `timeZones.txt`

// TimeZone represents a single time zone object
type TimeZone struct {
	CountryIso2Code string        // CountryCode
	TimeZoneID      string        // TimeZoneId
	GmtOffset       time.Duration // GMT offset 1. Jan 2016
	DstOffset       time.Duration // DST offset 1. Jul 2016
	RawOffset       time.Duration // rawOffset (independant of DST)
}

// TimeZones returns all time zones available
func (c *Client) TimeZones() ([]*TimeZone, error) {
	var err error
	var result []*TimeZone

	data, err := httpGet(geonamesURL + timeZonesURL)
	if err != nil {
		return nil, err
	}

	parse(data, 1, func(raw [][]byte) bool {
		if len(raw) != 5 {
			return true
		}

		gmtOffset, _ := strconv.ParseFloat(string(raw[2]), 64)
		dstOffset, _ := strconv.ParseFloat(string(raw[3]), 64)
		rawOffset, _ := strconv.ParseFloat(string(raw[4]), 64)

		result = append(result, &TimeZone{
			CountryIso2Code: string(raw[0]),
			TimeZoneID:      string(raw[1]),
			GmtOffset:       time.Duration(gmtOffset) * time.Hour,
			DstOffset:       time.Duration(dstOffset) * time.Hour,
			RawOffset:       time.Duration(rawOffset) * time.Hour,
		})

		return true
	})

	return result, nil
}
