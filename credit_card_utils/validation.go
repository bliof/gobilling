package credit_card_utils

func IsValidMonth(month int) bool {
	return month > 0 && month <= 12
}
