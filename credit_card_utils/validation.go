package credit_card_utils

import (
	"fmt"
	"time"
)

const (
	REQUIRED = "required"
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

func CheckForRequiredFields(rcc ReadOnlyCreditCard) error {
	e := new(FieldsError)

	if rcc.GetFirstName() == "" {
		e.Add("FirstName", REQUIRED)
	}

	if rcc.GetLastName() == "" {
		e.Add("LastName", REQUIRED)
	}

	if rcc.GetNumber() == "" {
		e.Add("Number", REQUIRED)
	}

	if rcc.GetBrand() == "" {
		e.Add("Brand", REQUIRED)
	}

	if rcc.GetMonth() == 0 {
		e.Add("Month", REQUIRED)
	}

	if rcc.GetYear() == 0 {
		e.Add("Year", REQUIRED)
	}

	return e.ToError()
}
