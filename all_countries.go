package geonames

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const allCountriesURI = `allCountries.zip`

// AllCountries returns a big pack of all features of all countries
func AllCountries() (map[int]*Feature, error) {
	var err error
	result := make(map[int]*Feature)
	url := fmt.Sprintf("%s%s", geonamesURL, allCountriesURI)

	dat, err := httpGetNew(url)
	if err != nil {
		return nil, err
	}

	tempPath := getTempPath(allCountriesURI)

	f, err := writeToFile(tempPath, dat)
	if err != nil {
		return nil, err
	}
	dat.Close()
	defer f.Close()
	defer os.Remove(tempPath)

	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return nil, err
	}
	defer r.Close()
	txtName := strings.Replace(allCountriesURI, "zip", "txt", -1)
	var s *bufio.Scanner
	for i := range r.File {
		if r.File[i].Name == txtName {
			rc, e := r.File[i].Open()
			if e != nil {
				return nil, e
			}
			defer rc.Close()

			s = bufio.NewScanner(rc)
			break
		}
	}

	if s == nil {
		return nil, errors.New("unknown error")
	}

	sParse(s, 0, func(raw []string) bool {
		if len(raw) != 19 {
			return true
		}

		geonameID, err := strconv.Atoi(raw[0])
		if err != nil {
			log.Printf("while parsing feature geoname id %s: %s", raw[0], err.Error())
			return true
		}

		alternateNames := strings.Split(raw[3], ",")
		for i := range alternateNames {
			alternateNames[i] = strings.TrimSpace(alternateNames[i])
			if alternateNames[i] == "" {
				alternateNames = append(alternateNames[:i], alternateNames[i+1:]...)
			}
		}

		latitude, err := strconv.ParseFloat(raw[4], 64)
		if err != nil {
			log.Printf("while parsing feature latitude %s: %s", raw[4], err.Error())
			return true
		}
		longitude, err := strconv.ParseFloat(raw[5], 64)
		if err != nil {
			log.Printf("while parsing feature longitude %s: %s", raw[5], err.Error())
			return true
		}

		var population *int
		if (raw[14]) != "" {
			populationInt, e := strconv.Atoi(raw[14])
			if e == nil {
				population = &populationInt
			}
		}

		var elevation *int
		if (raw[15]) != "" {
			elevationInt, e := strconv.Atoi(raw[15])
			if e == nil {
				elevation = &elevationInt
			}
		}

		dem, err := strconv.Atoi(raw[16])
		if err != nil {
			log.Printf("while parsing feature dem %s: %s", raw[16], err.Error())
			return true
		}
		modificationDate, err := time.Parse("2006-01-02", raw[18])
		if err != nil {
			log.Printf("while parsing feature modification date %s: %s", raw[18], err.Error())
			return true
		}

		result[geonameID] = &Feature{
			GeonameID:        geonameID,
			Name:             raw[1],
			ASCIIName:        raw[2],
			AlternateNames:   alternateNames,
			Latitude:         latitude,
			Longitude:        longitude,
			Class:            raw[6],
			Code:             raw[7],
			CountryCode:      raw[8],
			Cc2:              raw[9],
			Admin1Code:       raw[10],
			Admin2Code:       raw[11],
			Admin3Code:       raw[12],
			Admin4Code:       raw[13],
			Population:       population,
			Elevation:        elevation,
			Dem:              dem,
			TimeZone:         raw[17],
			ModificationDate: modificationDate,
		}

		return true
	})

	return result, nil
}
