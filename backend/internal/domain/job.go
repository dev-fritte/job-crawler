package domain

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Job struct {
	ExternalID      string
	Source          string
	URL             string
	Title           string
	Description     string
	Company         Company
	Language        Language
	Salary          *Salary
	RemoteMode      *RemoteMode
	Locations       []Location
	EmploymentTypes []EmploymentType
	Contract        ContractType
	SeniorityLevel  *SeniorityLevel
	PostedAt        *time.Time
	ValidThrough    *time.Time
	Tags            []string
}

func (j Job) Validate() error {

	var errs error

	if j.ExternalID == "" {
		errs = errors.Join(errs, errors.New("external ID is required"))
	}

	if j.Source == "" {
		errs = errors.Join(errs, errors.New("source is required"))
	}

	if j.URL == "" {
		errs = errors.Join(errs, errors.New("URL is required"))
	}

	if j.Title == "" {
		errs = errors.Join(errs, errors.New("title is required"))
	}

	if j.Company.Name == "" {
		errs = errors.Join(errs, errors.New("company is required"))
	}

	if j.Language != "" && !j.Language.Valid() {
		errs = errors.Join(errs, fmt.Errorf("invalid language: %s", j.Language))
	}

	if j.RemoteMode != nil && !j.RemoteMode.Valid() {
		errs = errors.Join(errs, fmt.Errorf("invalid remote mode: %s", *j.RemoteMode))
	}

	if j.SeniorityLevel != nil && !j.SeniorityLevel.Valid() {
		errs = errors.Join(errs, fmt.Errorf("invalid seniority level: %s", *j.SeniorityLevel))
	}

	if j.Salary != nil && !j.Salary.Currency.Valid() {
		errs = errors.Join(errs, fmt.Errorf("invalid salary currency: %s", j.Salary.Currency))
	}

	return errs
}

func (j *Job) AddTag(tag string) {

	tag = strings.ToLower(strings.TrimSpace(tag))

	if tag == "" {
		return
	}

	if j.HasTag(tag) {
		return
	}

	j.Tags = append(j.Tags, tag)
}

func (j Job) HasTag(tag string) bool {
	if j.Tags == nil {
		return false
	}

	tag = strings.ToLower(strings.TrimSpace(tag))

	return slices.Contains(j.Tags, tag)
}

func (s Salary) IsRange() bool {
	return s.Min != 0 && s.Max != 0 && s.Min != s.Max
}
