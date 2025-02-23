package service

type PalindromService interface {
	IsPalindrom(word string) string
}

type palindromService struct {
}

func NewPalindromService() PalindromService {
	return &palindromService{}
}

func (s *palindromService) CheckPalindrom(text string) bool {
	reverseText := ""

	for i := len(text) - 1; i >= 0; i-- {
		reverseText += string(text[i])
	}

	if text == reverseText {
		return true
	} else {
		return false
	}
}

func (s *palindromService) IsPalindrom(text string) string {

	isPlaindrom := s.CheckPalindrom(text)
	if !isPlaindrom {
		return "Not palindrome"
	}

	return "Palindrome"
}
