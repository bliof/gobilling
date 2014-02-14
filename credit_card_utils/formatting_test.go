package credit_card_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayNumber(t *testing.T) {
	maskedNumber := MaskNumber("42224242422222")

	assert.Equal(t, "XXXX-XXXX-XXXX-2222", maskedNumber)
}
