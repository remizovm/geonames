package geonames

import (
	"log"
	"strconv"
)

const alternateNamesURL = `alternateNames.zip`

// AlternateName represents a single feature's alternate name
type AlternateName struct {
	ID              int    // alternateNameId   : the id of this alternate name, int
	GeonameID       int    // geonameid         : geonameId referring to id in table 'geoname', int
	IsoLanguage     string // isolanguage       : iso 639 language code 2- or 3-characters; 4-characters 'post' for postal codes and 'iata','icao' and faac for airport codes, fr_1793 for French Revolution names,  abbr for abbreviation, link for a website, varchar(7)
	Name            string // alternate name    : alternate name or name variant, varchar(200)
	IsPreferredName bool   // isPreferredName   : '1', if this alternate name is an official/preferred name
	IsShortName     bool   // isShortName       : '1', if this is a short name like 'California' for 'State of California'
	IsColloquial    bool   // isColloquial      : '1', if this alternate name is a colloquial or slang term
	IsHistoric      bool   // isHistoric        : '1', if this alternate name is historic and was used in the past
}

// AlternateNames returns alternate names for all features available
func AlternateNames() ([]*AlternateName, error) {
	var err error
	var result []*AlternateName

	zipped, err := httpGet(geonamesURL + alternateNamesURL)
	if err != nil {
		return nil, err
	}

	files, err := unzip(zipped)
	if err != nil {
		return nil, err
	}

	data, err := getZipData(files, "alternateNames.txt")
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 8 {
			return true
		}

		if string(raw[2]) == "link" {
			return true
		}

		id, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			log.Printf("while converting alternate name %s modification id: %s", string(raw[0]), err.Error())
			return true
		}
		geonameID, err := strconv.Atoi(string(raw[1]))
		if err != nil {
			log.Printf("while converting alternate name %s modification geoname id: %s", string(raw[1]), err.Error())
			return true
		}

		result = append(result, &AlternateName{
			ID:              id,
			GeonameID:       geonameID,
			IsoLanguage:     string(raw[2]),
			Name:            string(raw[3]),
			IsPreferredName: string(raw[4]) == boolTrue,
			IsShortName:     string(raw[5]) == boolTrue,
			IsColloquial:    string(raw[6]) == boolTrue,
			IsHistoric:      string(raw[7]) == boolTrue,
		})

		return true
	})

	return result, nil
}
