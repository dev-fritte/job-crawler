package domain

import "slices"

type Language string

const (
	LanguageDE Language = "de"
	LanguageEN Language = "en"
)

var Languages = []Language{
	LanguageDE,
	LanguageEN,
}

func (l Language) Valid() bool {
	return slices.Contains(Languages, l)
}
