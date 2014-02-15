package credit_card_utils

type ReadOnlyCreditCard interface {
	GetFirstName() string
	GetLastName() string
	GetNumber() string
	GetBrand() string
	GetVerificationValue() string
	GetMonth() int
	GetYear() int
	GetStartMonth() int
	GetStartYear() int
	GetIssueNumber() string
	RequiresVerificationValue() bool
}

type CreditCard interface {
	GetFirstName() string
	SetFirstName(string)

	GetLastName() string
	SetLastName(string)

	GetNumber() string
	SetNumber(string)

	GetBrand() string
	SetBrand(string)

	GetVerificationValue() string
	SetVerificationValue(string)

	RequiresVerificationValue() bool

	GetMonth() int
	SetMonth(int)

	GetYear() int
	SetYear(int)

	GetStartMonth() int
	SetStartMonth(int)

	GetStartYear() int
	SetStartYear(int)

	GetIssueNumber() string
	SetIssueNumber(string)
}
