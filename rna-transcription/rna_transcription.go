// Package strand: given a DNA strand, returns its RNA complement
package strand

const testVersion = 3

// ToRNA
func ToRNA(s string) string {
	var out string
	for _, v := range s {
		switch v {
		case 71: // G
			out += "C"
		case 67: // C
			out += "G"
		case 84: // T
			out += "A"
		case 65: // A
			out += "U"
		}
	}
	return out
}
