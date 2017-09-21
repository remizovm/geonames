package geonames

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type FeatureCode struct {
	Code            string
	Description     string
	DescriptionLong string
}

func (c *Client) FeatureCodes(lang string) ([]FeatureCode, error) {
	uri := fmt.Sprintf("featureCodes_%s.txt", strings.ToLower(lang))
	data, err := httpGet(geonamesURL + uri)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = '\t'
	result := []FeatureCode{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		featureCode := FeatureCode{
			Code:            record[0],
			Description:     record[1],
			DescriptionLong: record[2],
		}
		result = append(result, featureCode)
	}
	return result, nil
}
