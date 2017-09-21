package geonames

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

const shapesAllLowURL = "shapes_all_low.zip"

type GeoJSON struct {
	Type         string      `json:"type"`
	Raw          interface{} `json:"coordinates"`
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

func decodeGeoJSON(obj *GeoJSON) error {
	switch obj.Type {
	case "Polygon":
		polygon, err := decodePolygon(obj.Raw)
		if err != nil {
			return err
		}
		obj.Polygon = polygon
		break
	case "MultiPolygon":
		multiPolygon, err := decodePolygonSet(obj.Raw)
		if err != nil {
			return err
		}
		obj.MultiPolygon = multiPolygon
		break
	default:
		log.Fatalf("unknown geometry type %s", obj.Type)
	}
	obj.Raw = nil // Get rid of processed raw coords to free up some memory
	return nil
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
		geoJSON := GeoJSON{}
		if err := json.Unmarshal(raw[1], &geoJSON); err != nil {
			log.Fatal(err)
		}
		if err := decodeGeoJSON(&geoJSON); err != nil {
			log.Fatalf("while decoding geoJSON: %s", err)
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
