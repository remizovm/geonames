package geonames

import (
	"fmt"
	"regexp"
	"strings"
)

const modificationsURL = `modifications-%d-%02d-%02d.txt`
const modificationsPattern = `(\d{1,7})\s(.+)`

func Modifications(year, month, day int) (map[string][]string, error) {
	uri := fmt.Sprintf(modificationsURL, year, month, day)

	data, err := httpGet(geonamesURL + uri)
	if err != nil {
		return nil, err
	}

	modificationsRe := regexp.MustCompile(modificationsPattern)
	matches := modificationsRe.FindAllStringSubmatch(string(data), -1)
	result := make(map[string][]string)

	for i := range matches {
		geonameId := matches[i][1]
		data := strings.Split(matches[i][2], "\t")
		result[geonameId] = data
	}

	return result, nil
}
