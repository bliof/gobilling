package gobilling

import ccutils "github.com/bliof/gobilling/credit_card_utils"

type CreditCard struct {
	FirstName, LastName, Number, Brand, VerificationValue string
	Month, Year                                           int

	StartMonth, StartYear int
	IssueNumber           string

	OptionalVerficationValue bool
}

func (cc *CreditCard) GetFirstName() string     { return cc.FirstName }
func (cc *CreditCard) SetFirstName(name string) { cc.FirstName = name }

func (cc *CreditCard) GetLastName() string     { return cc.LastName }
func (cc *CreditCard) SetLastName(name string) { cc.LastName = name }

func (cc *CreditCard) GetNumber() string       { return cc.Number }
func (cc *CreditCard) SetNumber(number string) { cc.Number = number }

func (cc *CreditCard) GetBrand() string      { return cc.Brand }
func (cc *CreditCard) SetBrand(brand string) { cc.Brand = brand }

func (cc *CreditCard) GetVerificationValue() string      { return cc.VerificationValue }
func (cc *CreditCard) SetVerificationValue(value string) { cc.VerificationValue = value }

func (cc *CreditCard) GetMonth() int      { return cc.Month }
func (cc *CreditCard) SetMonth(month int) { cc.Month = month }

func (cc *CreditCard) GetYear() int     { return cc.Year }
func (cc *CreditCard) SetYear(year int) { cc.Year = year }

func (cc *CreditCard) GetStartMonth() int      { return cc.StartMonth }
func (cc *CreditCard) SetStartMonth(month int) { cc.StartMonth = month }

func (cc *CreditCard) GetStartYear() int     { return cc.StartYear }
func (cc *CreditCard) SetStartYear(year int) { cc.StartYear = year }

func (cc *CreditCard) GetIssueNumber() string            { return cc.IssueNumber }
func (cc *CreditCard) SetIssueNumber(issueNumber string) { cc.IssueNumber = issueNumber }

func (cc *CreditCard) RequiresVerificationValue() bool { return !cc.OptionalVerficationValue }
func (cc *CreditCard) SetRequiresVerificationValue(required bool) {
	cc.OptionalVerficationValue = !required
}

func (cc *CreditCard) Validate() error {
	return ccutils.ValidateCreditCard(cc)
}

func (cc *CreditCard) DisplayNumber() string {
	return ccutils.MaskNumber(cc.Number)
}

type BillingAddress struct {
	Street, City, State, Zip string
}
