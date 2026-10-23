package domain

import "slices"

type Salary struct {
	Min      int
	Max      int
	Currency Currency
	Period   SalaryPeriod
}

type Currency string

const (
	CurrencyEUR Currency = "EUR"
	CurrencyCHF Currency = "CHF"
)

var currencies = []Currency{
	CurrencyEUR,
	CurrencyCHF,
}

func (c Currency) Valid() bool {
	return slices.Contains(currencies, c)
}

type SalaryPeriod string

const (
	SalaryPeriodHourly  SalaryPeriod = "hourly"
	SalaryPeriodMonthly SalaryPeriod = "monthly"
	SalaryPeriodYearly  SalaryPeriod = "yearly"
)

var salaryPeriods = []SalaryPeriod{
	SalaryPeriodHourly,
	SalaryPeriodMonthly,
	SalaryPeriodYearly,
}

func (p SalaryPeriod) Valid() bool {
	return slices.Contains(salaryPeriods, p)
}
