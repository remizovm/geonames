package geonames

import (
	"strconv"

	"github.com/remizovm/geonames/models"
)

const admin2CodesURL = `admin2Codes.txt`

// Admin2Codes returns all admin2 codes encoded in ASCII
func (c *Client) Admin2Codes() ([]*models.AdminCode, error) {
	var err error
	var result []*models.AdminCode

	data, err := httpGet(geonamesURL + admin2CodesURL)
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		geonameID, _ := strconv.ParseInt(string(raw[3]), 10, 64)

		result = append(result, &models.AdminCode{
			Codes:     string(raw[0]),
			Name:      string(raw[1]),
			ASCIIName: string(raw[2]),
			GeonameID: geonameID,
		})

		return true
	})

	return result, nil
}
