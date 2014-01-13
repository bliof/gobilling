package gobilling

type CreditCard struct {
	FirstName, LastName, Number, Type, Month, VerificationValue string
	Year                                                        int
}

func (cc *CreditCard) Validate() (err error) {
	return
}

func (cc *CreditCard) DisplayNumber() string {
	return "XXXX-XXXX-XXXX-something"
}
