package domain

import "slices"

type EmploymentType string

const (
	EmploymentTypeFullTime       EmploymentType = "full_time"
	EmploymentTypePartTime       EmploymentType = "part_time"
	EmploymentTypeWorkingStudent EmploymentType = "working_student"
	EmploymentTypeInternship     EmploymentType = "internship"
	EmploymentTypeApprenticeship EmploymentType = "apprenticeship"
	EmploymentTypeFreelance      EmploymentType = "freelance"
	EmploymentTypeOther          EmploymentType = "other"
)

var EmploymentTypes = []EmploymentType{
	EmploymentTypeFullTime,
	EmploymentTypePartTime,
	EmploymentTypeWorkingStudent,
	EmploymentTypeInternship,
	EmploymentTypeApprenticeship,
	EmploymentTypeFreelance,
	EmploymentTypeOther,
}

func (e EmploymentType) Valid() bool {
	return slices.Contains(EmploymentTypes, e)
}

type ContractType string

const (
	ContractTypePermanent ContractType = "permanent"
	ContractTypeFixedTerm ContractType = "fixed_term"
)

var ContractTypes = []ContractType{
	ContractTypePermanent,
	ContractTypeFixedTerm,
}

func (c ContractType) Valid() bool {
	return slices.Contains(ContractTypes, c)
}

type RemoteMode string

const (
	RemoteModeOnsite RemoteMode = "onsite"
	RemoteModeHybrid RemoteMode = "hybrid"
	RemoteModeRemote RemoteMode = "remote"
)

var RemoteModes = []RemoteMode{
	RemoteModeOnsite,
	RemoteModeHybrid,
	RemoteModeRemote,
}

func (r RemoteMode) Valid() bool {
	return slices.Contains(RemoteModes, r)
}

type SeniorityLevel string

const (
	SeniorityLevelJunior    SeniorityLevel = "junior"
	SeniorityLevelMid       SeniorityLevel = "mid"
	SeniorityLevelSenior    SeniorityLevel = "senior"
	SeniorityLevelLead      SeniorityLevel = "lead"
	SeniorityLevelPrincipal SeniorityLevel = "principal"
)

var SeniorityLevels = []SeniorityLevel{
	SeniorityLevelJunior,
	SeniorityLevelMid,
	SeniorityLevelSenior,
	SeniorityLevelLead,
	SeniorityLevelPrincipal,
}

func (s SeniorityLevel) Valid() bool {
	return slices.Contains(SeniorityLevels, s)
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
