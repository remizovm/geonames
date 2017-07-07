package models

// AlternateName represents a single feature's alternate name
type AlternateName struct {
	ID              int    // alternateNameId   : the id of this alternate name, int
	GeonameID       int    // geonameid         : geonameId referring to id in table 'geoname', int
	IsoLanguage     string // isolanguage       : iso 639 language code 2- or 3-characters; 4-characters 'post' for postal codes and 'iata','icao' and faac for airport codes, fr_1793 for French Revolution names,  abbr for abbreviation, link for a website, varchar(7)
	Name            string // alternate name    : alternate name or name variant, varchar(200)
	IsPreferredName bool   // isPreferredName   : '1', if this alternate name is an official/preferred name
	IsShortName     bool   // isShortName       : '1', if this is a short name like 'California' for 'State of California'
	IsColloquial    bool   // isColloquial      : '1', if this alternate name is a colloquial or slang term
	IsHistoric      bool   // isHistoric        : '1', if this alternate name is historic and was used in the past
}
