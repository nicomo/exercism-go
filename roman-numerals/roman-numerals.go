// Package romannumerals converts arab numerals to roman
package romannumerals

import (
	"errors"
	"strings"
)

const testVersion = 2

type arab2Roman struct {
	arab  int
	roman string
}

var a2r = []arab2Roman{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n >= 4000 {
		return "", errors.New("invalid number")
	}

	var res string
	var count int

	for i := len(a2r) - 1; i >= 0; i-- {
		count, n = divAndRemain(n, a2r[i].arab)
		res += strings.Repeat(a2r[i].roman, count)
	}

	return res, nil
}

func divAndRemain(n, d int) (int, int) {
	return n / d, n % d
}
