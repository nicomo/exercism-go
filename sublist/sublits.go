// Package sublist checks if a list is contained within another one
package sublist

import "reflect"

// Relation specifies the relationship between 2 lists
type Relation string

// Relations
const (
	RelEqual     Relation = "equal"
	RelSublist   Relation = "sublist"
	RelSuperlist Relation = "superlist"
	RelUnequal   Relation = "unequal"
)

// Sublist checks if a list is contained within another one
func Sublist(listOne, listTwo []int) Relation {
	if reflect.DeepEqual(listOne, listTwo) {
		return RelEqual
	} else if contains(listOne, listTwo) {
		return RelSuperlist
	} else if contains(listTwo, listOne) {
		return RelSublist
	}
	return RelUnequal
}

func contains(listA, listB []int) bool {
	if len(listB) == 0 {
		return true
	} else if len(listB) > len(listA) {
		return false
	}

	for i := 0; i <= len(listA)-len(listB); i++ {
		if listA[i] != listB[0] {
			continue
		}
		remaining := true
		for j, e := range listB {
			remaining = remaining && listA[i+j] == e
		}
		if remaining {
			return true
		}
	}
	return false

}
