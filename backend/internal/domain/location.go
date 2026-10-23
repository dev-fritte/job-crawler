package domain

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
	CountryDE = "DE"
	CountryAT = "AT"
	CountryCH = "CH"
)
