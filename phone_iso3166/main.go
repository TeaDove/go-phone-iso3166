package phone_iso3166

import "strconv"

// GetCountry
//
//	returns ISO3166-alfa2 country code from phoneNumber
//	returns InvalidPhoneNumber if phone number is invalid
//	i.e. GetCountry(79778725196) returns "RU"
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

// GetCountryFromString
//
//	phoneNumber should be without plus sign
//	returns ISO3166-alfa2 country code from phoneNumber
//	returns InvalidPhoneNumber if phone number is invalid
//	i.e. GetCountryFromString("79778725196") returns "RU"
func GetCountryFromString(phoneNumber string) (string, error) {
	phoneNumberInt, err := strconv.ParseInt(phoneNumber, 10, 64)
	if err != nil {
		return "", err
	}
	return GetCountry(uint64(phoneNumberInt))
}

// MustGetCountryFromString
//
//	same as GetCountryFromString, but raises panic on error
func MustGetCountryFromString(phoneNumber string) string {
	result, err := GetCountryFromString(phoneNumber)
	check(err)
	return result
}
