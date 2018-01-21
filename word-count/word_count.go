package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency {
	re := regexp.MustCompile("[[:punct:]]")
	f := Frequency{}
	w := strings.Fields(strings.ToLower(phrase))
	for _, v := range w {
		v = re.ReplaceAllString(v, "")
		if v != "" {
			f[v]++
		}

	}
	return f
}
