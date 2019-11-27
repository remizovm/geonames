package geonames

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	geonamesURL   = "http://download.geonames.org/export/dump/"
	commentSymbol = byte('#')
	newLineSymbol = byte('\n')
	delimSymbol   = byte('\t')
	boolTrue      = "1"
)

func getTempPath(name string) string {
	tempDir := os.TempDir()
	return path.Join(tempDir, name)
}

func writeToFile(fileName string, data io.ReadCloser) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, data)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *Client) getRaw(url, name string) (*bufio.Scanner, error) {
	var err error
	resp, err := c.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	lowName := strings.ToLower(name)

	isZip := false
	if strings.Contains(lowName, "zip") {
		isZip = true
	}

	tempDir := os.TempDir()

	filePath := path.Join(tempDir, name)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	defer os.Remove(filePath)

	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	if written != resp.ContentLength {
		errMsg := fmt.Sprintf("%s %d %d", file.Name(), written, resp.ContentLength)
		return nil, errors.New(errMsg)
	}

	var result *bufio.Scanner
	if isZip {
		r, e := zip.OpenReader(file.Name())
		if e != nil {
			return nil, e
		}
		defer r.Close()
		txtName := strings.Replace(name, "zip", "txt", -1)
		for i := range r.File {
			if r.File[i].Name == txtName {
				rc, e := r.File[i].Open()
				if e != nil {
					return nil, e
				}
				defer rc.Close()

				result = bufio.NewScanner(rc)
				break
			}
		}
	} else {
		result = bufio.NewScanner(file)
	}

	return result, nil
}

func (c *Client) httpGet(url string) ([]byte, error) {
	var err error
	resp, err := c.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) httpGetNew(url string) (io.ReadCloser, error) {
	var err error
	resp, err := c.c.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
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

func sParse(s *bufio.Scanner, headerLength uint, f func([]string) bool) {
	var err error
	var line string
	var rawSplit []string
	for s.Scan() {
		if headerLength != 0 {
			headerLength--
			continue
		}
		line = s.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == commentSymbol {
			continue
		}
		rawSplit = strings.Split(line, "\t")
		if !f(rawSplit) {
			break
		}
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(data []byte, headerLength int, f func([][]byte) bool) {
	rawSplit := bytes.Split(data, []byte{newLineSymbol})
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
		rawLineSplit = bytes.Split(rawSplit[i], []byte{delimSymbol})
		if !f(rawLineSplit) {
			break
		}
	}
}
