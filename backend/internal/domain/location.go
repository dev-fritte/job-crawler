package domain

import "slices"

type Location struct {
	Label      string
	PostalCode string
	City       string
	Region     string
	Country    Country
	Latitude   *float64
	Longitude  *float64
}

type Country string

var (
	CountryDE Country = "DE"
	CountryAT Country = "AT"
	CountryCH Country = "CH"
)

var Countries = []Country{
	CountryDE,
	CountryAT,
	CountryCH,
}

func (c Country) Valid() bool {
	return slices.Contains(Countries, c)
}
