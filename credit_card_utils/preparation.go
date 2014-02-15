package credit_card_utils

import (
	"regexp"
	"strings"
)

func SetupNumber(number string) string {
	r := regexp.MustCompile(`[^\d]+`)

	return r.ReplaceAllString(number, "")
}

func SetupCreditCard(cc CreditCard) {
	cc.SetNumber(SetupNumber(cc.GetNumber()))

	if cc.GetBrand() != "" {
		cc.SetBrand(strings.ToLower(cc.GetBrand()))
	} else {
		cc.SetBrand(DetectBrand(cc.GetNumber()))
	}
}

func DetectBrand(number string) string {
	detectors := BrandDetectors()

	for _, brand := range BrandsByDetectorPriority() {
		if detectors[brand](number) {
			return brand
		}
	}

	return ""
}

func BrandsByDetectorPriority() []string {
	return []string{
		"visa", "master", "discover", "american_express", "diners_club",
		"jcb", "switch", "solo", "dankort", "forbrugsforeningen",
		"laser", "maestro",
	}
}

func BrandDetectors() map[string]func(string) bool {
	return map[string]func(string) bool{
		"visa":               DetectVisa,
		"master":             DetectMasterCard,
		"discover":           DetectDiscover,
		"american_express":   DetectAmericanExpress,
		"diners_club":        DetectDinersClub,
		"jcb":                DetectJCB,
		"switch":             DetectSwitch,
		"solo":               DetectSolo,
		"dankort":            DetectDankort,
		"maestro":            DetectMaestro,
		"forbrugsforeningen": DetectForbrugsforeningen,
		"laser":              DetectLaser,
	}
}

func DetectVisa(number string) bool {
	re := regexp.MustCompile(`^4\d{12}(\d{3})?$`)

	return re.MatchString(number)
}

func DetectMasterCard(number string) bool {
	re := regexp.MustCompile(`^(5[1-5]\d{4}|677189)\d{10}$`)

	return re.MatchString(number)
}

func DetectDiscover(number string) bool {
	re := regexp.MustCompile(`^(6011|65\d{2}|64[4-9]\d)\d{12}|(62\d{14})$`)

	return re.MatchString(number)
}

func DetectAmericanExpress(number string) bool {
	re := regexp.MustCompile(`^3[47]\d{13}$`)

	return re.MatchString(number)
}

func DetectDinersClub(number string) bool {
	re := regexp.MustCompile(`^3(0[0-5]|[68]\d)\d{11}$`)

	return re.MatchString(number)
}

func DetectJCB(number string) bool {
	re := regexp.MustCompile(`^35(28|29|[3-8]\d)\d{12}$`)

	return re.MatchString(number)
}

func DetectSwitch(number string) bool {
	re := regexp.MustCompile(`^6759\d{12}(\d{2,3})?$`)

	return re.MatchString(number)
}

func DetectSolo(number string) bool {
	re := regexp.MustCompile(`^6767\d{12}(\d{2,3})?$`)

	return re.MatchString(number)
}

func DetectDankort(number string) bool {
	re := regexp.MustCompile(`^5019\d{12}$`)

	return re.MatchString(number)
}

func DetectMaestro(number string) bool {
	re := regexp.MustCompile(`^(5[06-8]|6\d)\d{10,17}$`)

	return re.MatchString(number)
}

func DetectForbrugsforeningen(number string) bool {
	re := regexp.MustCompile(`^600722\d{10}$`)

	return re.MatchString(number)
}

func DetectLaser(number string) bool {
	if len(number) >= 6 && number[0:6] == "677189" {
		return false
	}

	re := regexp.MustCompile(`^(6304|6706|6709|6771)\d{8}(\d{4}|\d{6,7})?$`)

	return re.MatchString(number)
}
