package geonames

import (
	"fmt"
	"log"
	"strconv"

	"github.com/remizovm/geonames/models"
)

const alternateNamesModificationsURL = `alternateNamesModifications-%d-%02d-%02d.txt`

// AlternateNamesModifications returns all alternate names modified at the selected date
func (c *Client) AlternateNamesModifications(year, month, day int) (map[int]*models.AlternateName, error) {
	var err error
	result := make(map[int]*models.AlternateName)

	uri := fmt.Sprintf(alternateNamesModificationsURL, year, month, day)

	data, err := c.httpGet(geonamesURL + uri)
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

		id, _ := strconv.Atoi(string(raw[0]))
		geonameID, err := strconv.Atoi(string(raw[1]))
		if err != nil {
			log.Printf("while converting alternate name %s modification id: %s", string(raw[0]), err.Error())
			return true
		}

		result[geonameID] = &models.AlternateName{
			ID:              id,
			GeonameID:       geonameID,
			IsoLanguage:     string(raw[2]),
			Name:            string(raw[3]),
			IsPreferredName: string(raw[4]) == boolTrue,
			IsShortName:     string(raw[5]) == boolTrue,
			IsColloquial:    string(raw[6]) == boolTrue,
			IsHistoric:      string(raw[7]) == boolTrue,
		}

		return true
	})

	return result, nil
}
