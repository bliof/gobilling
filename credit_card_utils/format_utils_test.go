package credit_card_utils

import (
	"testing"
)

func TestDisplayNumber(t *testing.T) {
	maskedNumber := MaskNumber("42224242422222")
	expected := "XXXX-XXXX-XXXX-2222"

	if maskedNumber !=  expected {
		t.Errorf("got %s expected %s", maskedNumber, expected)
	}
}
