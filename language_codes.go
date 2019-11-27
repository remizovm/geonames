package geonames

import "github.com/remizovm/geonames/models"

const languageCodesURL = `iso-languagecodes.txt`

// LanguageCodes returns all available languages
func (c *Client) LanguageCodes() ([]*models.LanguageCode, error) {
	var err error
	var result []*models.LanguageCode

	data, err := c.httpGet(geonamesURL + languageCodesURL)
	if err != nil {
		return nil, err
	}

	parse(data, 1, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		result = append(result, &models.LanguageCode{
			Iso3: string(raw[0]),
			Iso2: string(raw[1]),
			Iso:  string(raw[2]),
			Name: string(raw[3]),
		})

		return true
	})

	return result, nil
}
