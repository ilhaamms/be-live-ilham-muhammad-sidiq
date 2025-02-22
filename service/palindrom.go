package service

type PalindromService interface {
	IsPalindrom(word string) string
}

type palindromService struct {
}

func NewPalindromService() PalindromService {
	return &palindromService{}
}

func (s *palindromService) CheckPalindrom(word string) bool {
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

func (s *palindromService) IsPalindrom(word string) string {

	isPlaindrom := s.CheckPalindrom(word)
	if !isPlaindrom {
		return "Not palindrome"
	}

	return "Palindrome"
}
