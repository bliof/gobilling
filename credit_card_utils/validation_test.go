package credit_card_utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValidMonth(t *testing.T) {
	for i := 1; i <= 12; i += 1 {
		assert.True(t, IsValidMonth(i), "%d is valid month number", i)
	}

	assert.False(t, IsValidMonth(0), "0 is incorrect month number")
	assert.False(t, IsValidMonth(0), "13 is incorrect month number")
}
