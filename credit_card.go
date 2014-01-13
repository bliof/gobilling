package gobilling

import ccutils "github.com/bliof/gobilling/credit_card_utils"

type CreditCard struct {
	FirstName, LastName, Number, Brand, Month, VerificationValue string
	Year                                                         int
}

func (cc *CreditCard) Validate() (err error) {
	cc.Brand = "visa"
	return
}

func (cc *CreditCard) DisplayNumber() string {
	return ccutils.MaskNumber(cc.Number)
}
