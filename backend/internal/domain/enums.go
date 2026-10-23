package domain

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

type ContractType string

const (
	ContractTypePermanent ContractType = "permanent"
	ContractTypeFixedTerm ContractType = "fixed_term"
)

type RemoteMode string

const (
	RemoteModeOnsite RemoteMode = "onsite"
	RemoteModeHybrid RemoteMode = "hybrid"
	RemoteModeRemote RemoteMode = "remote"
)

type SeniorityLevel string

const (
	SeniorityLevelJunior    SeniorityLevel = "junior"
	SeniorityLevelMid       SeniorityLevel = "mid"
	SeniorityLevelSenior    SeniorityLevel = "senior"
	SeniorityLevelLead      SeniorityLevel = "lead"
	SeniorityLevelPrincipal SeniorityLevel = "principal"
)

type SalaryPeriod string

const (
	SalaryPeriodHourly  SalaryPeriod = "hourly"
	SalaryPeriodMonthly SalaryPeriod = "monthly"
	SalaryPeriodYearly  SalaryPeriod = "yearly"
)
