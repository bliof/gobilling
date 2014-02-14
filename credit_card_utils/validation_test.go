package credit_card_utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsValidMonth(t *testing.T) {
	for i := 1; i <= 12; i += 1 {
		assert.True(t, IsValidMonth(i), "%d is valid month number", i)
	}

	assert.False(t, IsValidMonth(0), "0 is incorrect month number")
	assert.False(t, IsValidMonth(0), "13 is incorrect month number")
}

func TestIsValidExpiryYear(t *testing.T) {
	assert.True(t, IsValidExpiryYear(time.Now().Year()))
	assert.True(t, IsValidExpiryYear(time.Now().Year()+10))

	assert.False(t, IsValidExpiryYear(time.Now().Year()-1))
}

func TestSetupNumber(t *testing.T) {
	assert.Equal(t, "1234567812345678", SetupNumber("1234-5678-1234-5678"))
	assert.Equal(t, "1234567812345678", SetupNumber("    1234   5678  1234 5678    "))
}
