package credit_card_utils

type ReadOnlyCreditCard interface {
	GetFirstName() string
	GetLastName() string
	GetNumber() string
	GetBrand() string
	GetVerificationValue() string
	GetMonth() int
	GetYear() int
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
}
