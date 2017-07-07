package models

// AdminCode represents a single admin code encoded in ASCII
type AdminCode struct {
	Codes     string
	Name      string
	ASCIIName string
	GeonameID int64
}
