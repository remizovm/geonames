package geonames

import (
	"encoding/json"
	"strconv"
)

const shapesJSONURL = "shapes_simplified_low.json.zip"

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		GeonameIDStr string `json:"geoNameId"`
	} `json:"properties"`
	Geometry GeoJSON `json:"geometry"`
}

func (c *Client) ShapesJSON() ([]Shape, error) {
	zipped, err := httpGet(geonamesURL + shapesJSONURL)
	if err != nil {
		return nil, err
	}
	f, err := unzip(zipped)
	if err != nil {
		return nil, err
	}
	data, err := getZipData(f, "shapes_simplified_low.json")
	if err != nil {
		return nil, err
	}
	result := []Shape{}
	var shapePack struct {
		Type        string    `json:"type"`
		FeatureList []Feature `json:"features"`
	}
	if err := json.Unmarshal(data, &shapePack); err != nil {
		return nil, err
	}
	for _, feature := range shapePack.FeatureList {
		geonameID, err := strconv.Atoi(feature.Properties.GeonameIDStr)
		if err != nil {
			return nil, err
		}
		if err := decodeGeoJSON(&feature.Geometry); err != nil {
			return nil, err
		}
		shape := Shape{
			GeonameID: geonameID,
			GeoJSON:   feature.Geometry,
		}
		result = append(result, shape)
	}
	return result, nil
}
