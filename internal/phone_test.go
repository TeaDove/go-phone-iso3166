package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_GetCountryRU_Ok(t *testing.T) {
	countryCode := MustGetCountry(79778725196)
	assert.Equal(t, countryCode, "RU")
}

func TestUnit_GetCountryKZ_Ok(t *testing.T) {
	countryCode := MustGetCountry(77051140675)
	assert.Equal(t, countryCode, "KZ")
}

func TestUnit_GetCountryUS_Ok(t *testing.T) {
	countryCode := MustGetCountry(16105579152)
	assert.Equal(t, countryCode, "US")
}

func TestUnit_GetCountryInvalidPhone_RaisesError(t *testing.T) {
	_, err := GetCountry(99008725196)
	assert.Error(t, err)
}
