package domain

import "slices"

type Salary struct {
	Min      *int
	Max      *int
	Currency Currency
	Period   SalaryPeriod
}

type Currency string

const (
	CurrencyEUR = "EUR"
	CurrencyCHF = "CHF"
)

var Currencies = []Currency{
	CurrencyEUR,
	CurrencyCHF,
}

func (c Currency) Valid() bool {
	return slices.Contains(Currencies, c)
}

type SalaryPeriod string

const (
	SalaryPeriodHourly  SalaryPeriod = "hourly"
	SalaryPeriodMonthly SalaryPeriod = "monthly"
	SalaryPeriodYearly  SalaryPeriod = "yearly"
)

var SalaryPeriods = []SalaryPeriod{
	SalaryPeriodHourly,
	SalaryPeriodMonthly,
	SalaryPeriodYearly,
}

func (p SalaryPeriod) Valid() bool {
	return slices.Contains(SalaryPeriods, p)
}
