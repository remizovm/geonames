package geonames

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	geonamesURL   = "http://download.geonames.org/export/dump/"
	commentSymbol = byte('#')
	newLineSymbol = byte('\n')
	boolTrue      = "1"
)

func httpGet(url string) ([]byte, error) {
	var err error
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func unzip(data []byte) ([]*zip.File, error) {
	var err error

	r, err := zip.NewReader(bytes.NewReader(data), (int64)(len(data)))
	if err != nil {
		return nil, err
	}

	return r.File, nil
}

func getZipData(files []*zip.File, name string) ([]byte, error) {
	var result []byte

	for _, f := range files {
		if f.Name == name {
			src, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer src.Close()

			result, err = ioutil.ReadAll(src)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func parse(data []byte, headerLength int, f func([][]byte) bool) {
	rawSplit := bytes.Split(data, []byte{'\n'})
	var rawLineSplit [][]byte
	for i := range rawSplit {
		if headerLength != 0 {
			headerLength--
			continue
		}
		if len(rawSplit[i]) == 0 {
			continue
		}
		if rawSplit[i][0] == commentSymbol {
			continue
		}
		rawLineSplit = bytes.Split(rawSplit[i], []byte{'\t'})
		if !f(rawLineSplit) {
			break
		}
	}
}
