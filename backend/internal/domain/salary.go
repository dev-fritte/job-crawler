package domain

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

type Salary struct {
	Min      int
	Max      int
	Currency Currency
	Period   SalaryPeriod
}

const (
	monthPerYear = 12
)

// errors
var (
	ErrInvalidCurrency = errors.New("the defined currency is not valid")
	ErrInvalidPeriod   = errors.New("the defined salary period is not valid")
	ErrInvalidRange    = errors.New("salary range cannot be negative")
)

func (s Salary) YearlyEUR(chfToEUR float64) (minValue, maxValue int, err error) {
	if !s.Currency.Valid() {
		return 0, 0, fmt.Errorf("currency %q: %w", s.Currency, ErrInvalidCurrency)
	}

	if s.Min < 0 || s.Max < 0 {
		return 0, 0, fmt.Errorf("salary range min: %v max: %v invalid: %w", s.Min, s.Max, ErrInvalidRange)
	}

	if chfToEUR <= 0 {
		return 0, 0, errors.New("invalid CHF to EUR exchange rate")
	}

	var factor float64

	switch s.Period {
	case SalaryPeriodMonthly:
		factor = monthPerYear
	case SalaryPeriodYearly:
		factor = 1
	default:
		return 0, 0, fmt.Errorf("period %q: %w", s.Period, ErrInvalidPeriod)
	}

	if s.Currency == CurrencyCHF {
		factor *= chfToEUR
	}

	return int(math.Round(float64(s.Min) * factor)), int(math.Round(float64(s.Max) * factor)), nil
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
	SalaryPeriodMonthly SalaryPeriod = "monthly"
	SalaryPeriodYearly  SalaryPeriod = "yearly"
)

var salaryPeriods = []SalaryPeriod{
	SalaryPeriodMonthly,
	SalaryPeriodYearly,
}

func (p SalaryPeriod) Valid() bool {
	return slices.Contains(salaryPeriods, p)
}

func (s Salary) IsRange() bool {
	return s.Min != 0 && s.Max != 0 && s.Min != s.Max
}
