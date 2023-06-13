# Phone ISO3166

[![GoDoc](https://godoc.org/github.com/teadove/go-phone-iso3166?status.svg)](https://godoc.org/github.com/teadove/go-phone-iso3166)

Small project to map an E.164 (international) phone number to the
ISO-3166-1 alpha 2 (two letter) country code, associated with that number.

GO version of same python project: [phone_iso3166](https://pypi.org/project/phone-iso3166/)



## Installing

```
go get -u github.com/teadove/go-phone-iso3166
```

## Example

```go
countryCode, _ := GetCountry("+79778725196")
print(countryCode) // prints "RU"
```


## Contact

Peter Ibragimov [@teadove](http://t.me/teadove)

## License

Match source code is available under the MIT [License](/LICENSE).
