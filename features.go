package geonames

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/remizovm/geonames/models"
)

// Features returns all features for the specified country iso2 code
func (c *Client) Features(iso2code string) ([]*models.Feature, error) {
	var err error
	var result []*models.Feature

	if len(iso2code) != 2 {
		return nil, errors.New("Invalid iso2code")
	}

	uri := fmt.Sprintf("%s%s.zip", geonamesURL, strings.ToUpper(iso2code))
	zipped, err := c.httpGet(uri)
	if err != nil {
		return nil, err
	}

	f, err := unzip(zipped)
	if err != nil {
		return nil, err
	}

	data, err := getZipData(f, strings.ToUpper(iso2code)+".txt")
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 19 {
			return true
		}

		geonameID, _ := strconv.Atoi(string(raw[0]))

		alternateNames := strings.Split(string(raw[3]), ",")
		for i := range alternateNames {
			alternateNames[i] = strings.TrimSpace(alternateNames[i])
			if alternateNames[i] == "" {
				alternateNames = append(alternateNames[:i], alternateNames[i+1:]...)
			}
		}

		latitude, _ := strconv.ParseFloat(string(raw[4]), 64)
		longitude, _ := strconv.ParseFloat(string(raw[5]), 64)

		var population *int
		if string(raw[14]) != "" {
			populationInt, err := strconv.Atoi(string(raw[14]))
			if err == nil {
				population = &populationInt
			}
		}

		var elevation *int
		if string(raw[15]) != "" {
			elevationInt, err := strconv.Atoi(string(raw[15]))
			if err == nil {
				elevation = &elevationInt
			}
		}

		dem, _ := strconv.Atoi(string(raw[16]))
		modificationDate, _ := time.Parse("2006-02-01", string(raw[18]))

		result = append(result, &models.Feature{
			GeonameID:        geonameID,
			Name:             string(raw[1]),
			ASCIIName:        string(raw[2]),
			AlternateNames:   alternateNames,
			Latitude:         latitude,
			Longitude:        longitude,
			Class:            string(raw[6]),
			Code:             string(raw[7]),
			CountryCode:      string(raw[8]),
			Cc2:              string(raw[9]),
			Admin1Code:       string(raw[10]),
			Admin2Code:       string(raw[11]),
			Admin3Code:       string(raw[12]),
			Admin4Code:       string(raw[13]),
			Population:       population,
			Elevation:        elevation,
			Dem:              dem,
			TimeZone:         string(raw[17]),
			ModificationDate: modificationDate,
		})

		return true
	})

	return result, nil
}
