package geonames

import (
	"fmt"
	"log"
	"strconv"
)

const deletesURL = `deletes-%d-%02d-%02d.txt`

// DeleteOp represents a single object deletion operation
type DeleteOp struct {
	GeonameID int
	Name      string
	Comment   string
}

// Deletes returns all deleted objects for the selected date
func Deletes(year, month, day int) (map[int]*DeleteOp, error) {
	var err error
	uri := fmt.Sprintf(deletesURL, year, month, day)

	data, err := httpGet(geonamesURL + uri)
	if err != nil {
		return nil, err
	}

	result := make(map[int]*DeleteOp)
	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 3 {
			return true
		}
		geonameID, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			log.Printf("while converting raw deletion %s geoname id: %s", string(raw[0]), err.Error())
			log.Println(string(raw[0]))
			return true
		}

		result[geonameID] = &DeleteOp{
			Name:    string(raw[1]),
			Comment: string(raw[2])}
		return true
	})

	return result, nil
}
