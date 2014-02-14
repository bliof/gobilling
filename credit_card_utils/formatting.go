package credit_card_utils

import "fmt"

func MaskNumber(number string) string {
	return fmt.Sprintf("XXXX-XXXX-XXXX-%s", number[len(number)-4:])
}
