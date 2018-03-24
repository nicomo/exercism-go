// Package brackets : given a string containing brackets `[]`, braces `{}` and parentheses `()`,
// verify that all the pairs are matched and nested correctly.
package brackets

type bracketType int

const (
	openingB bracketType = iota
	closingB
	notB
)

var pairs = map[rune]rune{'[': ']', '{': '}', '(': ')'}

// Bracket ensures that opening brackets close properly
func Bracket(s string) (bool, error) {
	var stack []rune
	for _, c := range s {
		switch getBracketType(c) {
		case openingB:
			stack = append(stack, pairs[c])
		case closingB:
			if 0 < len(stack) && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
				continue
			}
			return false, nil
		}

	}

	return len(stack) == 0, nil
}

func getBracketType(r rune) bracketType {
	for k, v := range pairs {
		switch r {
		case k:
			return openingB
		case v:
			return closingB
		default:
		}
	}
	return notB
}
