package gateway

import "github.com/bliof/gobilling"

type PayPal struct {
	User, Password, Signature string
}

func (pp *PayPal) Purchase(amount float64, creditCard gobilling.CreditCard) (response *Response) {
	return
}

type Response struct{}

func (r *Response) IsSuccessful() bool {
	return true
}
