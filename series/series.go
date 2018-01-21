// package slice will take a string of digits and give you all the contiguous substrings of length n in that string.
// For example, the string "49142" has the following 3-digit series:
// 491
// 914
// 142
package slice

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var sls []string
	for i := range s {
		if i+n <= len(s) {
			sls = append(sls, s[i:i+n])
		}
	}
	return sls
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First signals that in some cases you can't take the first N characters of the string
func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		ok = true
		first = s[0:n]
	}
	return first, ok
}
