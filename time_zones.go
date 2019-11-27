package geonames

import (
	"strconv"
	"time"

	"github.com/remizovm/geonames/models"
)

const timeZonesURL = `timeZones.txt`

// TimeZones returns all time zones available
func (c *Client) TimeZones() ([]*models.TimeZone, error) {
	var err error
	var result []*models.TimeZone

	data, err := c.httpGet(geonamesURL + timeZonesURL)
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

		result = append(result, &models.TimeZone{
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
