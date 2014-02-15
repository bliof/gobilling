package gateway

import (
	"testing"
	"time"

	"github.com/bliof/gobilling"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/stretchr/testify/assert"
)

func TestPayPalPurchase(t *testing.T) {
	config, _ := yaml.ReadFile("test_config.yaml")

	if config == nil {
		return
	}

	user, _ := config.Get("paypal.user")
	password, _ := config.Get("paypal.password")
	signature, _ := config.Get("paypal.signature")

	if user == "" || password == "" || signature == "" {
		return
	}

	gateway := PayPal{
		User:      user,
		Password:  password,
		Signature: signature,
		Testing:   true,
	}

	creditCard := gobilling.CreditCard{
		FirstName:         "Rose",
		LastName:          "Tyler",
		Number:            "4660997962602322",
		Month:             9,
		Year:              time.Now().Year() + 1,
		VerificationValue: "176",
	}

	err := creditCard.Validate()

	if err == nil {
		response := gateway.Purchase(
			100,
			&creditCard,
			gobilling.BillingAddress{Street: "FirstStreet", City: "SanJose", State: "CA", Zip: "95131"},
		)

		assert.True(t, response.IsSuccessful())
	}
}
