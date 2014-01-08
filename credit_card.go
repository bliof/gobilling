package gobilling

type CreditCard struct{
	FirstName, LastName, Number, Type, Month, VerificationValue string
	Year int
}

func (cc *CreditCard) Validate() (valid bool, err error) {
	return
}
