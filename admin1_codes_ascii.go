package geonames

import "strconv"

const admin1CodesAsciiURL = `admin1CodesASCII.txt`

type Admin1CodeAscii struct {
	Codes     string
	Name      string
	AsciiName string
	GeonameID int64
}

func Admin1CodesAscii() ([]*Admin1CodeAscii, error) {
	var err error
	var result []*Admin1CodeAscii

	data, err := httpGet(geonamesURL + admin1CodesAsciiURL)
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		geonameID, _ := strconv.ParseInt(string(raw[3]), 10, 64)

		result = append(result, &Admin1CodeAscii{
			Codes:     string(raw[0]),
			Name:      string(raw[1]),
			AsciiName: string(raw[2]),
			GeonameID: geonameID,
		})

		return true
	})

	return result, nil
}
