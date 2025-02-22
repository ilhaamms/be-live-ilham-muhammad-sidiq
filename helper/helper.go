package helper

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
