// package dna
package dna

import "errors"

// Dna is a DNA representation
type Dna struct {
	histogram Histogram
}

// Histogram is a count of each type of nucleotide in a dna strand
type Histogram map[byte]int

// DNA gets the number of valid bytes in a string
// weeds out invalid nucleaotides
func DNA(s string) Dna {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range s {
		h[byte(n)]++
	}
	return Dna{h}
}

// Count retrieves the number of occurences
// of a particular nucleotides in a strand of DNA
func (dna Dna) Count(n byte) (int, error) {
	count, ok := dna.histogram[n]
	err := errors.New("invalid nucleotide")
	if !ok {
		return 0, err
	}
	return count, nil
}

// Counts returns a histogram
// i.e. number of occurences of all nucleotides
// in a given DNA strand
func (dna Dna) Counts() Histogram {
	return dna.histogram
}
