package geonames

import (
	"fmt"
	"log"
	"strconv"

	"github.com/remizovm/geonames/models"
)

const deletesURL = `deletes-%d-%02d-%02d.txt`

// Deletes returns all deleted objects for the selected date
func (c *Client) Deletes(year, month, day int) (map[int]*models.DeleteOp, error) {
	var err error
	uri := fmt.Sprintf(deletesURL, year, month, day)

	data, err := c.httpGet(geonamesURL + uri)
	if err != nil {
		return nil, err
	}

	result := make(map[int]*models.DeleteOp)
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

		result[geonameID] = &models.DeleteOp{
			Name:    string(raw[1]),
			Comment: string(raw[2])}
		return true
	})

	return result, nil
}
