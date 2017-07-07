package models

// AlternateNameDeleteOp represents a single operation of feature's AlternateName deletion
type AlternateNameDeleteOp struct {
	ID        int
	GeonameID int
	Name      string
	Comment   string
}
