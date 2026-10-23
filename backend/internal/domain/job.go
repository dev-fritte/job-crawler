package domain

type Job struct {
	ID          string
	Title       string
	Description string
	Company     Company
	URL         string
	Source      string
	ExternalID  string
	Language    Language
}
