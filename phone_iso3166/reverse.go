package phone_iso3166

import (
	"math/bits"
)

const reversedPhoneNumberCap = 12

type reversedPhoneNumber []byte

func reverseNumber(num uint64) reversedPhoneNumber {
	res := make(reversedPhoneNumber, 0, reversedPhoneNumberCap)
	var rem uint64
	for num > 0 {
		num, rem = bits.Div64(0, num, 10)
		// TODO possibly very slow, maybe optimize?
		res = append(reversedPhoneNumber{byte(rem)}, res...)
	}
	return res
}
