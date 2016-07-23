package geonames

import (
	"log"
	"strconv"
)

const hierarchyURL = `hierarchy.zip`

type HierarchyNode struct {
	ParentID int
	ChildID  int
	Type     string
}

func Hierarchy() (map[int][]*HierarchyNode, error) {
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
		parentId, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			log.Printf("while parsing hierarchy parent id %s: %s", string(raw[0]), err.Error())
			return true
		}
		childId, err := strconv.Atoi(string(raw[1]))
		if err != nil {
			log.Printf("while parsing hierarchy child id %s: %s", string(raw[1]), err.Error())
			return true
		}

		result[parentId] = append(result[parentId], &HierarchyNode{
			ParentID: parentId,
			ChildID:  childId,
			Type:     string(raw[2])})

		return true
	})

	return result, nil
}
