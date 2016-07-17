package geonames

import "strconv"

const admin2CodesURL = `admin2Codes.txt`

type Admin2Code struct {
	Codes     string
	Name      string
	AsciiName string
	GeonameID int64
}

func Admin2Codes() ([]*Admin2Code, error) {
	var err error
	var result []*Admin2Code

	data, err := httpGet(geonamesURL + admin2CodesURL)
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		geonameID, _ := strconv.ParseInt(string(raw[3]), 10, 64)

		result = append(result, &Admin2Code{
			Codes:     string(raw[0]),
			Name:      string(raw[1]),
			AsciiName: string(raw[2]),
			GeonameID: geonameID,
		})

		return true
	})

	return result, nil
}
