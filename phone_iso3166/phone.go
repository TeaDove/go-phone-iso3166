package phone_iso3166

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
