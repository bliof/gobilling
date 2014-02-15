package credit_card_utils

import (
	"fmt"
	"regexp"
	"time"
)

const (
	REQUIRED = "required"
	INVALID  = "invalid"
)

type FieldError struct {
	Field string
	Type  string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s is %s", e.Field, e.Type)
}

type FieldsError struct {
	FieldErrors []FieldError
	Message     string
}

func (e *FieldsError) Error() string {
	message := e.Message

	for i, err := range e.FieldErrors {
		message += fmt.Sprintf("%d) %s", i, err.Error())
	}

	return message
}

func (e *FieldsError) Add(field string, errorType string) {
	e.FieldErrors = append(e.FieldErrors, FieldError{Field: field, Type: errorType})
}

func (e *FieldsError) ToError() error {
	if len(e.FieldErrors) != 0 || e.Message != "" {
		return e
	}

	return nil
}

func IsValidMonth(month int) bool {
	return month > 0 && month <= 12
}

func IsValidExpiryYear(year int) bool {
	currentYear := time.Now().Year()
	return year >= currentYear && year <= currentYear+20
}

func IsValidStartYear(year int) bool {
	return year > 1987
}

func IsValidIssueNumber(number string) bool {
	re := regexp.MustCompile(`^\d{1,2}$`)

	return re.MatchString(number)
}

func CheckForRequiredFields(rcc ReadOnlyCreditCard) error {
	err := new(FieldsError)

	if rcc.GetFirstName() == "" {
		err.Add("FirstName", REQUIRED)
	}

	if rcc.GetLastName() == "" {
		err.Add("LastName", REQUIRED)
	}

	if rcc.GetBrand() == "" {
		err.Add("Brand", REQUIRED)
	}

	if rcc.GetNumber() == "" {
		err.Add("Number", REQUIRED)
	}

	if rcc.RequiresVerificationValue() && rcc.GetVerificationValue() == "" {
		err.Add("VerificationValue", REQUIRED)
	}

	if rcc.GetMonth() == 0 {
		err.Add("Month", REQUIRED)
	}

	if rcc.GetYear() == 0 {
		err.Add("Year", REQUIRED)
	}

	return err.ToError()
}

func ValidateCreditCard(cc CreditCard) error {
	SetupCreditCard(cc)

	required := CheckForRequiredFields(cc)

	if required != nil {
		return required
	}

	err := new(FieldsError)

	if !IsValidMonth(cc.GetMonth()) {
		err.Add("Month", INVALID)
	}

	if !IsValidExpiryYear(cc.GetYear()) {
		err.Add("Year", INVALID)
	}

	if cc.GetBrand() == "switch" || cc.GetBrand() == "solo" {
		if !(IsValidMonth(cc.GetStartMonth()) && IsValidStartYear(cc.GetStartYear()) || IsValidIssueNumber(cc.GetIssueNumber())) {
			if cc.GetIssueNumber() == "" {
				err.Add("IssueNumber", REQUIRED)

				if !IsValidMonth(cc.GetStartMonth()) {
					err.Add("StartMonth", INVALID)
				}

				if !IsValidStartYear(cc.GetStartYear()) {
					err.Add("StartYear", INVALID)
				}
			} else {
				if !IsValidIssueNumber(cc.GetIssueNumber()) {
					err.Add("IssueNumber", INVALID)
				}
			}
		}
	}

	return err.ToError()
}
