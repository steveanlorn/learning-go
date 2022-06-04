// Package caesarchipher provides mechanism to substitute cipher
// in which each letter in the plaintext is replaced by a letter
// some fixed number of positions down the alphabet.
package caesarchipher

import "strings"

// Encrypt converts into a ciphertext of the given string
// by replacing a letter by shifting the position
// to the right in the alphabet.
func Encrypt(s string, shift uint) string {
	return convert("right", s, shift)
}

// Decrypt converts a ciphertext into original text of the given string
// by replacing a letter by shifting the position
// to the left in the alphabet.
func Decrypt(s string, shift uint) string {
	return convert("left", s, shift)
}

func convert(direction string, s string, shift uint) string {
	var plainText string
	s = strings.ToLower(s)

	for _, r := range s {
		shiftedRune := shiftRune(direction, r, shift)
		plainText += string(shiftedRune)
	}

	return plainText
}

func shiftRune(direction string, r rune, s uint) rune {
	shift := rune(s % 26)
	if shift == 0 {
		return r
	}

	var shiftedRune rune

	if direction == "right" {
		shiftedRune = r + shift
	} else {
		shiftedRune = r - shift
	}

	if shiftedRune < 'a' {
		shiftedRune = shiftedRune + 26
	} else if shiftedRune > 'z' {
		shiftedRune = shiftedRune - 26
	}

	return shiftedRune
}
