package geonames

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Feature struct {
	GeonameID        int       // geonameid         : integer id of record in geonames database
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
	Population       *int      // population        : bigint (8 byte int)
	Elevation        *int      // elevation         : in meters, integer
	Dem              int       // dem               : digital elevation model, srtm3 or gtopo30, average elevation of 3''x3'' (ca 90mx90m) or 30''x30'' (ca 900mx900m) area in meters, integer. srtm processed by cgiar/ciat.
	TimeZone         string    // timezone          : the timezone id (see file timeZone.txt) varchar(40)
	ModificationDate time.Time // modification date : date of last modification in yyyy-MM-dd format
}

func Features(iso2code string) ([]*Feature, error) {
	var err error
	var result []*Feature

	if len(iso2code) != 2 {
		return nil, errors.New("Invalid iso2code")
	}

	uri := fmt.Sprintf("%s%s.zip", geonamesURL, strings.ToUpper(iso2code))
	zipped, err := httpGet(uri)
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

		geonameId, _ := strconv.Atoi(string(raw[0]))

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

		result = append(result, &Feature{
			GeonameID:        geonameId,
			Name:             string(raw[1]),
			AsciiName:        string(raw[2]),
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
