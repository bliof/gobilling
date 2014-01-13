package gobilling

import (
	"testing"
	"time"
)

func creditCard() CreditCard {
	return CreditCard{
		FirstName:         "Rose",
		LastName:          "Tyler",
		Number:            "4222222222222",
		Month:             "9",
		Year:              time.Now().Year() + 1,
		VerificationValue: "000",
	}
}

func TestDisplayNumber(t *testing.T) {
	cc := creditCard()

	if cc.DisplayNumber() != "XXXX-XXXX-XXXX-2222" {
		t.Error(cc.DisplayNumber())
	}
}
