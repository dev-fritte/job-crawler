package domain

import (
	"fmt"
	"slices"
	"strings"
)

type EmploymentType string

const (
	EmploymentTypeFullTime       EmploymentType = "full_time"
	EmploymentTypePartTime       EmploymentType = "part_time"
	EmploymentTypeWorkingStudent EmploymentType = "working_student"
	EmploymentTypeInternship     EmploymentType = "internship"
	EmploymentTypeApprenticeship EmploymentType = "apprenticeship"
	EmploymentTypeFreelance      EmploymentType = "freelance"
)

var employmentTypes = []EmploymentType{
	EmploymentTypeFullTime,
	EmploymentTypePartTime,
	EmploymentTypeWorkingStudent,
	EmploymentTypeInternship,
	EmploymentTypeApprenticeship,
	EmploymentTypeFreelance,
}

func (e EmploymentType) Valid() bool {
	return slices.Contains(employmentTypes, e)
}

type ContractType string

const (
	ContractTypePermanent ContractType = "permanent"
	ContractTypeFixedTerm ContractType = "fixed_term"
)

var contractTypes = []ContractType{
	ContractTypePermanent,
	ContractTypeFixedTerm,
}

func (c ContractType) Valid() bool {
	return slices.Contains(contractTypes, c)
}

type RemoteMode string

const (
	RemoteModeOnsite RemoteMode = "onsite"
	RemoteModeHybrid RemoteMode = "hybrid"
	RemoteModeRemote RemoteMode = "remote"
)

var remoteModes = []RemoteMode{
	RemoteModeOnsite,
	RemoteModeHybrid,
	RemoteModeRemote,
}

func (r RemoteMode) Valid() bool {
	return slices.Contains(remoteModes, r)
}

var remoteAliases = map[string]RemoteMode{
	"on-site":    RemoteModeOnsite,
	"on site":    RemoteModeOnsite,
	"vor ort":    RemoteModeOnsite,
	"hybride":    RemoteModeHybrid,
	"homeoffice": RemoteModeRemote,
}

func ParseRemoteMode(s string) (RemoteMode, error) {
	normalizedString := strings.ToLower(strings.TrimSpace(s))

	if remoteMode := RemoteMode(normalizedString); remoteMode.Valid() {
		return remoteMode, nil
	}

	if mode, ok := remoteAliases[normalizedString]; ok {
		return mode, nil
	}

	return "", fmt.Errorf("invalid remote mode: %q", s)
}

type SeniorityLevel string

const (
	SeniorityLevelJunior    SeniorityLevel = "junior"
	SeniorityLevelMid       SeniorityLevel = "mid"
	SeniorityLevelSenior    SeniorityLevel = "senior"
	SeniorityLevelLead      SeniorityLevel = "lead"
	SeniorityLevelPrincipal SeniorityLevel = "principal"
)

var seniorityLevels = []SeniorityLevel{
	SeniorityLevelJunior,
	SeniorityLevelMid,
	SeniorityLevelSenior,
	SeniorityLevelLead,
	SeniorityLevelPrincipal,
}

func (s SeniorityLevel) Valid() bool {
	return slices.Contains(seniorityLevels, s)
}
