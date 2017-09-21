package geonames

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

const shapesAllLowURL = "shapes_all_low.zip"

type GeoJSON struct {
	Type         string `json:"type"`
	Polygon      [][][]float64
	MultiPolygon [][][][]float64
}

type Shape struct {
	GeonameID int
	GeoJSON   GeoJSON
}

func decodePoint(data interface{}) ([]float64, error) {
	coords, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("not a valid set of points")
	}
	result := make([]float64, 0, len(coords))
	for _, coord := range coords {
		f, ok := coord.(float64)
		if !ok {
			return nil, errors.New("got invalid point")
		}
		result = append(result, f)
	}
	return result, nil
}

func decodePositionSet(data interface{}) ([][]float64, error) {
	points, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("not a valid set of points")
	}
	result := make([][]float64, 0, len(points))
	for _, point := range points {
		p, err := decodePoint(point)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func decodePolygonSet(data interface{}) ([][][][]float64, error) {
	polygonList, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("not a valid multipolygon data")
	}
	result := make([][][][]float64, 0, len(polygonList))
	for _, polygon := range polygonList {
		p, err := decodePolygon(polygon)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil

}

func decodePolygon(data interface{}) ([][][]float64, error) {
	sets, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("not a valid polygon data")
	}
	result := make([][][]float64, 0, len(sets))
	for _, set := range sets {
		s, err := decodePositionSet(set)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (c *Client) Shapes() ([]Shape, error) {
	zipped, err := httpGet(geonamesURL + shapesAllLowURL)
	if err != nil {
		return nil, err
	}
	f, err := unzip(zipped)
	if err != nil {
		return nil, err
	}
	data, err := getZipData(f, "shapes_all_low.txt")
	if err != nil {
		return nil, err
	}
	result := []Shape{}
	parse(data, 1, func(raw [][]byte) bool {
		geonameID, _ := strconv.Atoi(string(raw[0]))
		var rawGeoJSON struct {
			Type   string      `json:"type"`
			Coords interface{} `json:"coordinates"`
		}
		if err := json.Unmarshal(raw[1], &rawGeoJSON); err != nil {
			panic(err)
		}
		geoJSON := GeoJSON{
			Type: rawGeoJSON.Type,
		}
		switch rawGeoJSON.Type {
		case "Polygon":
			polygon, err := decodePolygon(rawGeoJSON.Coords)
			if err != nil {
				panic(err)
			}
			geoJSON.Polygon = polygon
			break
		case "MultiPolygon":
			multiPolygon, err := decodePolygonSet(rawGeoJSON.Coords)
			if err != nil {
				panic(err)
			}
			geoJSON.MultiPolygon = multiPolygon
			break
		default:
			log.Fatalf("unknown geometry type %s", rawGeoJSON.Type)
		}
		shape := Shape{
			GeonameID: geonameID,
			GeoJSON:   geoJSON,
		}
		result = append(result, shape)
		return true
	})
	return result, nil
}
