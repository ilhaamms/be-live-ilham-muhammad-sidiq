package helper

// function untuk mengecek apakah kata palindrom atau tidak
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
