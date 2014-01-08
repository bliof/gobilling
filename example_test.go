package gobilling_test

import (
	"fmt"
	"time"

	"github.com/bliof/gobilling"
)

// Here we create a credit card object and validate.
// Notice that when you validate the card the Type will be
// automaticly added to the struct if it was not provided in
// advance.
func ExampleCreditCard_create() {
	// The verificationValue is also known as CVV2, CVC2, or CID
	creditCard := gobilling.CreditCard{
		FirstName: "Rose",
		LastName: "Tyler",
		Number: "4242424242424242",
		Month: "9",
		Year: time.Now().Year() + 1,
		VerificationValue: "000",
	}

	// When validating the credit card, the type will be automaticly filled
	ok, _ := creditCard.Validate()

	fmt.Printf("Is the card valid? %t", ok)
	fmt.Printf("creditCard.Type = %s", creditCard.Type)
	// Output:
	// Is the card valid? true
	// creditCard.Type = visa
}
