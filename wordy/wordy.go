// Package wordy takes a sentence problem and calculates result
package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

// Answer turns string into int operation
func Answer(s string) (int, bool) {
	var myints []int
	var mysigns []string
	var res int
	re := regexp.MustCompile("\\?")

	w := strings.Fields(s)
	for _, v := range w {
		v = re.ReplaceAllString(v, "")
		if myint, err := strconv.Atoi(v); err == nil {
			myints = append(myints, myint)
		} else {
			switch v {
			case "plus":
				mysigns = append(mysigns, "+")
			case "minus":
				mysigns = append(mysigns, "-")
			case "divided":
				mysigns = append(mysigns, "/")
			case "multiplied":
				mysigns = append(mysigns, "*")
			}
		}
	}

	if len(mysigns) == 0 {
		return 0, false
	}

	for i := 0; i < len(mysigns); i++ {
		switch mysigns[i] {
		case "+":
			if res == 0 {
				res = myints[i] + myints[i+1]
			} else {
				res = res + myints[i+1]
			}

		case "-":
			if res == 0 {
				res = myints[i] - myints[i+1]
			} else {
				res = res - myints[i+1]
			}

		case "/":
			if res == 0 {
				res = myints[i] / myints[i+1]
			} else {
				res = res / myints[i+1]
			}

		case "*":
			if res == 0 {
				res = myints[i] * myints[i+1]
			} else {
				res = res * myints[i+1]
			}
		}

	}

	return res, true

}
