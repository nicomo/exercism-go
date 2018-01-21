// Package kindergarten matches kids with plants
// You must define a type Garden with constructor
//
//    func NewGarden(diagram string, children []string) (*Garden, error)
//
// and method
//
//    func (g *Garden) Plants(child string) ([]string, bool)
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

const testVersion = 1

//plants maps initials to the full names of the plants
var plants = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// Garden represents the cups on a window sill
type Garden map[string][]string

// NewGarden creates a garden
func NewGarden(diagram string, children []string) (*Garden, error) {

	// sort the children
	sortedC := make([]string, len(children))
	copy(sortedC, children)
	sort.Strings(sortedC)

	// check for duplicates in the children list
	for k, v := range sortedC[1:] {
		if sortedC[k] == v {
			return nil, errors.New("Duplicate kid")
		}
	}

	d := strings.Split(diagram, "\n")

	err := diagramValidate(d, len(sortedC))
	if err != nil {
		return nil, err
	}

	// create garden with just the children first
	g := Garden{}
	for _, c := range sortedC {
		g[c] = make([]string, 0)
	}

	// look at plants in rows
	for _, row := range d {
		if len(row) < 2 {
			continue
		}
		for idx, r := range row {
			plant, ok := plants[r]
			if !ok {
				return nil, errors.New("unknown plant")
			}
			currentList := g[sortedC[idx/2]]
			g[sortedC[idx/2]] = append(currentList, plant)
		}
	}

	return &g, nil
}

// Plants retrieves the list of plants attributed in the Garden to a given child
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}

func diagramValidate(d []string, l int) error {

	// garden should have 2 rows
	if len(d)-1 != 2 {
		return errors.New("wrong number of rows in garden")
	}

	// 2 rows should have the same length
	if len(d[1]) != len(d[2]) {
		return errors.New("mismatched rows")
	}

	// number of cups should be even
	if len(d[1])%2 != 0 {
		return errors.New("odd number of cups")
	}

	// garden should have 2 plants per row per children
	if l != len(d[1])/2 || l != len(d[2])/2 {
		return errors.New("wrong number of plants in row")
	}

	return nil

}
