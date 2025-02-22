package helper

import "github.com/go-playground/validator/v10"

func IsPalindrom(word string) bool {
	reverseWord := ""

	for i := len(word) - 1; i >= 0; i-- {
		reverseWord += string(word[i])
	}

	if word == reverseWord {
		return true
	} else {
		return false
	}
}

var Validate = validator.New()
