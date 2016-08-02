package geonames

import (
	"fmt"
	"log"
	"strconv"
)

const alternateNamesdeletesURL = `alternateNamesdeletes-%d-%02d-%02d.txt`

// AlternateNameDeleteOp represents a single operation of feature's AlternateName deletion
type AlternateNameDeleteOp struct {
	ID        int
	GeonameID int
	Name      string
	Comment   string
}

// AlternateNamesDeletes returns all deleted alternate names for the selected date
func AlternateNamesDeletes(year, month, day int) (map[int]*AlternateNameDeleteOp, error) {
	var err error
	uri := fmt.Sprintf(alternateNamesdeletesURL, year, month, day)

	data, err := httpGet(geonamesURL + uri)
	if err != nil {
		return nil, err
	}

	result := make(map[int]*AlternateNameDeleteOp)
	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 4 {
			return true
		}

		id, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			log.Printf("while converting alternate name deletion %s id: %s", string(raw[0]), err.Error())
			return true
		}

		geonameID, err := strconv.Atoi(string(raw[1]))
		if err != nil {
			log.Printf("while converting alternate name deletion %s geoname id: %s", string(raw[1]), err.Error())
			return true
		}

		result[geonameID] = &AlternateNameDeleteOp{
			ID:        id,
			GeonameID: geonameID,
			Name:      string(raw[2]),
			Comment:   string(raw[3])}
		return true
	})

	return result, nil
}
