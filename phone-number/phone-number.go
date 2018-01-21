// Package phonenumber cleans up user-entered phone numbers
package phonenumber

import (
	"errors"
	"unicode"
)

// Number: is a phone number valid or not?
// if it is, return it sanitized
func Number(in string) (string, error) {
	var a string

	for _, v := range in {
		if unicode.IsDigit(v) {
			a += string(v)
		}
	}

	if len(a) < 10 || len(a) > 11 {
		err := errors.New("too short, or too long")
		return "", err
	}

	if len(a) == 11 {
		if a[:1] == "1" {
			b := a[1:]
			return b, nil
		} else {
			err := errors.New("11 but doesn't start with 1")
			return "", err
		}
	}

	if len(a) == 10 {
		return a, nil
	}

	return "", nil
}

// AreaCode extracts the 1st 3 digits from a valid number
func AreaCode(in string) (string, error) {
	s, err := Number(in)
	if err != nil {
		return "", err
	}

	return s[:3], nil

}

// Format formats a valid string into a phone number
func Format(in string) (string, error) {
	s, err := Number(in)
	if err != nil {
		return "", err
	}
	out := "(" + s[:3] + ") " + s[3:6] + "-" + s[6:]
	return out, nil
}
