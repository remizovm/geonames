package models

// HierarchyNode represents a pair of parent and child objects linked together
type HierarchyNode struct {
	ParentID int
	ChildID  int
	Type     string
}
