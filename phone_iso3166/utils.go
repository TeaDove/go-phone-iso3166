package phone_iso3166

import "log"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
