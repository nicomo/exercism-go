// Package protein translates RNA sequences into proteins
package protein

const testVersion = 1

var codonsMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns the protein or Stop for a given codon
func FromCodon(s string) string {
	return codonsMap[s]
}

// FromRNA returns the proteins for a RNA strand
func FromRNA(s string) []string {
	var r []string

	for len(s) > 0 {
		if FromCodon(s[:3]) == "STOP" {
			return r
		}
		r = append(r, FromCodon(s[:3]))
		s = s[3:]
	}
	return r
}
