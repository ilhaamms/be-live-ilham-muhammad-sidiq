package service

import (
	"github.com/ilhaamms/backend-live/helper"
)

type PalindromService interface {
	IsPalindrom(word string) string
}

type palindromService struct {
}

func NewPalindromService() PalindromService {
	return &palindromService{}
}

func (service *palindromService) IsPalindrom(word string) string {

	isPlaindrom := helper.IsPalindrom(word)
	if !isPlaindrom {
		return "Not palindrome"
	}

	return "Palindrome"
}
