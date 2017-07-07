package models

// Country represents a single country
type Country struct {
	Iso2Code           string  // ISO
	Iso3Code           string  // ISO3
	IsoNumeric         string  // ISO-Numeric
	Fips               string  // fips
	Name               string  // Country
	Capital            string  // Capital
	Area               float64 // Area(in sq km)
	Population         uint64  // Population
	Continent          string  // Continent
	Tld                string  // tld
	CurrencyCode       string  // CurrencyCode
	CurrencyName       string  // CurrencyName
	Phone              string  // Phone
	PostalCodeFormat   string  // Postal Code Format
	PostalCodeRegex    string  // Postal Code Regex
	Languages          string  // Languages
	GeonameID          int64   // geonameid
	Neighbours         string  // neighbours
	EquivalentFipsCode string  // EquivalentFipsCode
}
