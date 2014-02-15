package gateway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bliof/gobilling"
)

type CreditCard interface {
	GetFirstName() string
	GetLastName() string
	GetNumber() string
	GetBrand() string
	GetVerificationValue() string
	GetMonth() int
	GetYear() int
}

type PayPal struct {
	User, Password, Signature string
	Testing                   bool
}

func (pp *PayPal) Version() string {
	return "78"
}

func (pp *PayPal) endpoint() string {
	if pp.Testing {
		return "https://api-3t.sandbox.paypal.com/nvp"
	} else {
		return "https://api-3t.paypal.com/nvp"
	}
}

func (pp *PayPal) Purchase(amount float64, cc CreditCard, ba gobilling.BillingAddress) Response {
	params := url.Values{}

	params.Add("USER", pp.User)
	params.Add("PWD", pp.Password)
	params.Add("SIGNATURE", pp.Signature)
	params.Add("METHOD", "DoDirectPayment")
	params.Add("VERSION", pp.Version())
	params.Add("PAYMENTACTION", "SALE")
	params.Add("AMT", fmt.Sprintf("%0.02f", amount))
	params.Add("ACCT", cc.GetNumber())
	params.Add("CREDITCARDTYPE", "VISA")
	params.Add("CVV2", cc.GetVerificationValue())
	params.Add("FIRSTNAME", cc.GetFirstName())
	params.Add("LASTNAME", cc.GetLastName())
	params.Add("EXPDATE", fmt.Sprintf("%02d%d", cc.GetMonth(), cc.GetYear()))
	params.Add("STREET", ba.Street)
	params.Add("City", ba.City)
	params.Add("State", ba.State)
	params.Add("ZIP", ba.Zip)

	response, err := http.Get(pp.endpoint() + "?" + params.Encode())

	if err != nil {
		return Response{successful: false, Error: err}
	}

	return parseResponse(response)
}

func parseResponse(response *http.Response) Response {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return Response{successful: false, Error: err}
	}

	params, params_err := url.ParseQuery(string(body))

	if params_err != nil {
		return Response{successful: false, Error: params_err}
	}

	return Response{successful: params["ACK"][0] == "Success", Params: params}
}

type Response struct {
	successful bool
	Error      error
	Params     url.Values
}

func (r *Response) IsSuccessful() bool {
	return r.successful
}
