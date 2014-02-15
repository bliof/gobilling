package gobilling_test

import (
	"fmt"
	"time"

	"github.com/bliof/gobilling"
	"github.com/bliof/gobilling/gateway"
)

// Here we create a credit card object and validate it.
// Notice that when you validate the card the Brand will be
// automaticly added to the struct if it was not provided in
// advance.
func ExampleCreditCard_create() {
	// The verificationValue is also known as CVV2, CVC2, or CID
	creditCard := gobilling.CreditCard{
		FirstName:         "Rose",
		LastName:          "Tyler",
		Number:            "4242424242424242",
		Month:             9,
		Year:              time.Now().Year() + 1,
		VerificationValue: "000",
	}

	// When validating the credit card, the Brand will be automaticly filled
	err := creditCard.Validate()

	fmt.Printf("Is the card valid? %t\n", err == nil)
	fmt.Printf("creditCard.Brand = %s\n", creditCard.Brand)
	// Output:
	// Is the card valid? true
	// creditCard.Brand = visa
}

func ExampleGateway_interaction() {
	gateway := gateway.PayPal{
		User:      "TestMerchant",
		Password:  "password",
		Signature: "ashjdfasdkf",
		Testing:   true,
	}

	amount := 20.0

	creditCard := gobilling.CreditCard{
		FirstName:         "Rose",
		LastName:          "Tyler",
		Number:            "4222222222222",
		Month:             9,
		Year:              time.Now().Year() + 1,
		VerificationValue: "000",
	}

	err := creditCard.Validate()

	if err == nil {
		response := gateway.Purchase(
			amount,
			&creditCard,
			gobilling.BillingAddress{Street: "FirstStreet", City: "SanJose", State: "CA", Zip: "95131"},
		)

		if response.IsSuccessful() {
			fmt.Printf("Charged %.2f to the credit card %s", amount, creditCard.DisplayNumber())
		}
	}
}
