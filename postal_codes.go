package geonames

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/remizovm/geonames/models"
)

const postalCodesURL = `http://download.geonames.org/export/zip/%s.zip`

// PostalCodes returns all postal codes for the selected countries iso2 code
func (c *Client) PostalCodes(iso2code string) (map[string]*models.PostalCode, error) {
	var err error
	result := make(map[string]*models.PostalCode)

	if len(iso2code) != 2 {
		return nil, errors.New("invalid iso2code")
	}

	url := fmt.Sprintf(postalCodesURL, strings.ToUpper(iso2code))
	zipped, err := httpGet(url)
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
		if len(raw) != 12 {
			return true
		}

		latitude, err := strconv.ParseFloat(string(raw[9]), 64)
		if err != nil {
			log.Printf("while parsing postal code latitude: %s", string(raw[9]))
			return true
		}
		longitude, err := strconv.ParseFloat(string(raw[10]), 64)
		if err != nil {
			log.Printf("while parsing postal code longitude: %s", string(raw[10]))
			return true
		}
		accuracy, err := strconv.Atoi(string(raw[11]))
		if err != nil {
			log.Printf("while parsing postal code accuracy: %s", string(raw[11]))
			return true
		}

		result[string(raw[1])] = &models.PostalCode{
			CountryIso2Code: string(raw[0]),
			PostalCode:      string(raw[1]),
			PlaceName:       string(raw[2]),
			AdminName1:      string(raw[3]),
			AdminCode1:      string(raw[4]),
			AdminName2:      string(raw[5]),
			AdminCode2:      string(raw[6]),
			AdminName3:      string(raw[7]),
			AdminCode3:      string(raw[8]),
			Latitude:        latitude,
			Longitude:       longitude,
			Accuracy:        accuracy,
		}

		return true
	})

	return result, nil
}
