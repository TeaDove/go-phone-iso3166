package phone_iso3166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ReverseNumber_Ok(t *testing.T) {
	reversed := reverseNumber(79778725196)

	assert.Equal(t, reversed, reversedPhoneNumber{7, 9, 7, 7, 8, 7, 2, 5, 1, 9, 6})
}

func TestUnit_ReverseNumberWithTrailingZero_Ok(t *testing.T) {
	reversed := reverseNumber(79778725100)

	assert.Equal(t, reversed, reversedPhoneNumber{7, 9, 7, 7, 8, 7, 2, 5, 1, 0, 0})
}
