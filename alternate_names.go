package geonames

import (
	"log"
	"strconv"

	"github.com/remizovm/geonames/models"
)

const alternateNamesURL = `alternateNames.zip`

// AlternateNames returns alternate names for all features available
func (c *Client) AlternateNames() ([]*models.AlternateName, error) {
	var err error
	var result []*models.AlternateName

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

		result = append(result, &models.AlternateName{
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
