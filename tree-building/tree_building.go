// Package tree refactors a tree
package tree

import (
	"errors"
	"sort"
)

// Record is the input struct
type Record struct {
	ID, Parent int
}

// Node is the recursive output struct
type Node struct {
	ID       int
	Children []*Node
}

// Build creates the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var result Node

	// order records by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	err := checks(records)
	if err != nil {
		return nil, err
	}

	for _, r := range records {

		// root record: is correct, not duplicate
		if r.ID == 0 {
			result.ID = r.ID
			continue
		}

		// children of root
		if result.ID == r.Parent {
			if result.Children == nil {
				result.Children = make([]*Node, 0)
			}
			child := Node{
				ID: r.ID,
			}
			result.Children = append(result.Children, &child)
			continue
		}

		// other children
		for _, n := range result.Children {
			if r.Parent == n.ID {
				child := Node{
					ID: r.ID,
				}
				n.Children = append(n.Children, &child)
			}
		}
	}

	return &result, nil
}

func checks(records []Record) error {
	ids := make(map[int]bool)
	for i, r := range records {
		if _, ok := ids[r.ID]; ok {
			return errors.New("nope, duplicate and such")
		}
		ids[r.ID] = true

		if (r.ID < r.Parent) ||
			(r.ID != 0 && r.Parent == r.ID) ||
			(i == 0 && r.ID != 0) ||
			(i != r.ID) {
			return errors.New("no no no")
		}
	}
	return nil
}
