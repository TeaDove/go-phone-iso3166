package internal

import (
	"errors"
)

var (
	InvalidPhoneNumber = errors.New("invalid phone number")
)

func traverse(phoneNumber reversedPhoneNumber, rune mapRune) (string, error) {
	firstNum := phoneNumber[0]
	remainingNumbers := phoneNumber[1:]

	nextRuneI, ok := rune[firstNum]
	if !ok {
		return "", errors.Join(errors.New("index out of bounds"), InvalidPhoneNumber)
	}
	country, ok := nextRuneI.(string)
	if ok {
		return country, nil
	}
	nextRune, ok := nextRuneI.(mapRune)
	if !ok {
		return "", errors.Join(errors.New("invalid rune loaded"), InvalidPhoneNumber)
	}
	return traverse(remainingNumbers, nextRune)
}

// GetCountry
//
//	returns ISO3166-alfa2 country code from phoneNumber
//	returns InvalidPhoneNumber if phone number is invalid
func GetCountry(phoneNumber uint64) (string, error) {
	return traverse(reverseNumber(phoneNumber), mapping)
}

// MustGetCountry
//
//	same as GetCountry, but raises panic on error
func MustGetCountry(phoneNumber uint64) string {
	country, err := GetCountry(phoneNumber)
	check(err)
	return country
}
