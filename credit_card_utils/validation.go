package credit_card_utils

import "time"

func IsValidMonth(month int) bool {
	return month > 0 && month <= 12
}

func IsValidExpiryYear(year int) bool {
	currentYear := time.Now().Year()
	return year >= currentYear && year <= currentYear+20
}
