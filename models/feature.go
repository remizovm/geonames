package models

import "time"

// Feature represents a single administrative object - for example, a city
type Feature struct {
	GeonameID        int       // geonameid         : integer id of record in geonames database
	Name             string    // name              : name of geographical point (utf8) varchar(200)
	ASCIIName        string    // asciiname         : name of geographical point in plain ascii characters, varchar(200)
	AlternateNames   []string  // alternatenames    : alternatenames, comma separated, ascii names automatically transliterated, convenience attribute from alternatename table, varchar(10000)
	Latitude         float64   // latitude          : latitude in decimal degrees (wgs84)
	Longitude        float64   // longitude         : longitude in decimal degrees (wgs84)
	Class            string    // feature class     : see http://www.geonames.org/export/codes.html, char(1)
	Code             string    // feature code      : see http://www.geonames.org/export/codes.html, varchar(10)
	CountryCode      string    // country code      : ISO-3166 2-letter country code, 2 characters
	Cc2              string    // cc2               : alternate country codes, comma separated, ISO-3166 2-letter country code, 200 characters
	Admin1Code       string    // admin1 code       : fipscode (subject to change to iso code), see exceptions below, see file admin1Codes.txt for display names of this code; varchar(20)
	Admin2Code       string    // admin2 code       : code for the second administrative division, a county in the US, see file admin2Codes.txt; varchar(80)
	Admin3Code       string    // admin3 code       : code for third level administrative division, varchar(20)
	Admin4Code       string    // admin4 code       : code for fourth level administrative division, varchar(20)
	Population       *int      // population        : bigint (8 byte int)
	Elevation        *int      // elevation         : in meters, integer
	Dem              int       // dem               : digital elevation model, srtm3 or gtopo30, average elevation of 3''x3'' (ca 90mx90m) or 30''x30'' (ca 900mx900m) area in meters, integer. srtm processed by cgiar/ciat.
	TimeZone         string    // timezone          : the timezone id (see file timeZone.txt) varchar(40)
	ModificationDate time.Time // modification date : date of last modification in yyyy-MM-dd format
}
