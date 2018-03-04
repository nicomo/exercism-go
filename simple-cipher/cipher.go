// Package cipher implements a simple shift cipher like Caesar and a more secure substitution cipher
package cipher

import (
	"strings"
	"unicode"
)

// Cipher is an encoding / decoding interface
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type cipherKey []int

// Encode encodes a plain text (pt) given a cipher key
func (c cipherKey) Encode(pt string) string {
	return c.EncodeDecode(pt, true)
}

// Decode decodes a ciphered text (ct) given a cipher key
func (c cipherKey) Decode(ct string) string {
	return c.EncodeDecode(ct, false)
}

// EncodeDecode encodes or decodes a string given a key and and order (encode == true)
func (c cipherKey) EncodeDecode(source string, encode bool) string {

	var r string
	shift := 0
	for _, letter := range source {
		var char int

		if !unicode.IsLetter(letter) {
			continue
		}

		if encode {
			char = int(unicode.ToLower(letter)) + c[shift]
		} else {
			char = int(unicode.ToLower(letter)) - c[shift]
		}

		// if letter + cipher is beyond 'z' in table, start again at 'a'
		switch {
		case char < 'a':
			char += 'z' - 'a' + 1
		case 'z' < char:
			char -= 'z' - 'a' + 1
		}

		r += string(char)

		if shift < len(c)-1 {
			shift++
		} else {
			shift = 0
		}

	}
	return r
}

// NewCaesar returns a new cipher with fixed shift of distance 3
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a cipherKey with a given fixed distance
func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	cKey := cipherKey([]int{distance})
	return cKey
}

// NewVigenere creates a Vigenère cipher
func NewVigenere(key string) Cipher {
	// validate key
	if strings.Count(key, "a") == len(key) || len(key) == 0 {
		return nil
	}

	vigCipher := make([]int, 0)

	for _, v := range key {
		if !unicode.IsLower(v) || !unicode.IsLetter(v) {
			return nil
		}
		vigCipher = append(vigCipher, int(v)-'a')
	}

	cKey := cipherKey(vigCipher)
	return cKey
}
