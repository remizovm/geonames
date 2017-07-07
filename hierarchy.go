package geonames

import (
	"log"
	"strconv"
)

const hierarchyURL = `hierarchy.zip`

// HierarchyNode represents a pair of parent and child objects linked together
type HierarchyNode struct {
	ParentID int
	ChildID  int
	Type     string
}

// Hierarchy returns all available pairs of linked parents and children
// For example, a city is linked to it's country as a child:
// Country->City1,City2 etc
func (c *Client) Hierarchy() (map[int][]*HierarchyNode, error) {
	var err error
	result := make(map[int][]*HierarchyNode)

	zipped, err := httpGet(geonamesURL + hierarchyURL)
	if err != nil {
		return nil, err
	}

	f, err := unzip(zipped)
	if err != nil {
		return nil, err
	}

	data, err := getZipData(f, "hierarchy.txt")
	if err != nil {
		return nil, err
	}

	parse(data, 0, func(raw [][]byte) bool {
		if len(raw) != 3 {
			return true
		}
		parentID, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			log.Printf("while parsing hierarchy parent id %s: %s", string(raw[0]), err.Error())
			return true
		}
		childID, err := strconv.Atoi(string(raw[1]))
		if err != nil {
			log.Printf("while parsing hierarchy child id %s: %s", string(raw[1]), err.Error())
			return true
		}

		result[parentID] = append(result[parentID], &HierarchyNode{
			ParentID: parentID,
			ChildID:  childID,
			Type:     string(raw[2])})

		return true
	})

	return result, nil
}
