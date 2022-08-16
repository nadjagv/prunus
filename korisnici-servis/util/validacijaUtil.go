package util

import (
	"regexp"
	"unicode"
)

func ValidanEmail(input string) bool {
	regex := "^[a-zA-Z0-9_+&*-]+(?:\\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,7}$"

	match, _ := regexp.MatchString(regex, input)
	return match
}

func ValidnaLozinka(input string) bool {
	slova := 0
	var broj, veliko, specijalni, barOsam bool
	for _, c := range input {
		switch {
		case unicode.IsNumber(c):
			broj = true
			slova++
		case unicode.IsUpper(c):
			veliko = true
			slova++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			specijalni = true
		case unicode.IsLetter(c) || c == ' ':
			slova++
		default:
		}
	}
	barOsam = slova >= 8

	if broj && veliko && specijalni && barOsam {
		return true
	}
	return false
}
