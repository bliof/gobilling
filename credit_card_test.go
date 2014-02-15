package gobilling

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func creditCard() CreditCard {
	return CreditCard{
		FirstName:         "Rose",
		LastName:          "Tyler",
		Number:            "4222222222222",
		Month:             9,
		Year:              time.Now().Year() + 1,
		VerificationValue: "000",
	}
}

func TestDisplayNumber(t *testing.T) {
	cc := creditCard()

	assert.Equal(t, "XXXX-XXXX-XXXX-2222", cc.DisplayNumber())
}

func TestValidate(t *testing.T) {
	cc := creditCard()
	err := cc.Validate()

	assert.Nil(t, err)
	assert.Equal(t, "visa", cc.Brand)
}
