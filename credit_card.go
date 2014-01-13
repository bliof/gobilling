package gobilling

type CreditCard struct {
	FirstName, LastName, Number, Brand, Month, VerificationValue string
	Year                                                         int
}

func (cc *CreditCard) Validate() (err error) {
	return
}

func (cc *CreditCard) DisplayNumber() string {
	return "XXXX-XXXX-XXXX-something"
}
