package geonames

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const postalCodesUrl = `http://download.geonames.org/export/zip/%s.zip`

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

func PostalCodes(iso2code string) (map[string]*PostalCode, error) {
	var err error
	result := make(map[string]*PostalCode)

	if len(iso2code) != 2 {
		return nil, errors.New("invalid iso2code")
	}

	url := fmt.Sprintf(postalCodesUrl, strings.ToUpper(iso2code))
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

		result[string(raw[1])] = &PostalCode{
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
