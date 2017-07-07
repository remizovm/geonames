package geonames

import "strconv"

const admin1CodesASCIIURL = `admin1CodesASCII.txt`

// Admin1CodeASCII represents a single admin1 code encoded in ASCII
type Admin1CodeASCII struct {
	Codes     string
	Name      string
	ASCIIName string
	GeonameID int64
}

// Admin1CodesASCII returns all admin1 codes encoded in ASCII
func (c *Client) Admin1CodesASCII() ([]*Admin1CodeASCII, error) {
	var err error
	var result []*Admin1CodeASCII

	data, err := httpGet(geonamesURL + admin1CodesASCIIURL)
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		geonameID, _ := strconv.ParseInt(string(raw[3]), 10, 64)

		result = append(result, &Admin1CodeASCII{
			Codes:     string(raw[0]),
			Name:      string(raw[1]),
			ASCIIName: string(raw[2]),
			GeonameID: geonameID,
		})

		return true
	})

	return result, nil
}
