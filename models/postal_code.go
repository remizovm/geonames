package models

// PostalCode represents a single countries postal code
type PostalCode struct {
	CountryIso2Code string  // country code      : iso country code, 2 characters
	PostalCode      string  // postal code       : varchar(20)
	PlaceName       string  // place name        : varchar(180)
	AdminName1      string  // admin name1       : 1. order subdivision (state) varchar(100)
	AdminCode1      string  // admin code1       : 1. order subdivision (state) varchar(20)
	AdminName2      string  // admin name2       : 2. order subdivision (county/province) varchar(100)
	AdminCode2      string  // admin code2       : 2. order subdivision (county/province) varchar(20)
	AdminName3      string  // admin name3       : 3. order subdivision (community) varchar(100)
	AdminCode3      string  // admin code3       : 3. order subdivision (community) varchar(20)
	Latitude        float64 // latitude          : estimated latitude (wgs84)
	Longitude       float64 // longitude         : estimated longitude (wgs84)
	Accuracy        int     // accuracy          : accuracy of lat/lng from 1=estimated to 6=centroi}
}
