package credit_card_utils

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupNumber(t *testing.T) {
	assert.Equal(t, "1234567812345678", SetupNumber("1234-5678-1234-5678"))
	assert.Equal(t, "1234567812345678", SetupNumber("    1234   5678  1234 5678    "))
}

func caller() string {
	_, file, line, _ := runtime.Caller(2)

	parts := strings.Split(file, "/")
	file = parts[len(parts)-1]

	return fmt.Sprintf("%s:%d", file, line)
}

func assertBrand(t *testing.T, brand string, number string) {
	assert.Equal(t, brand, DetectBrand(number), "%s Should detect %s for %s", caller(), brand, number)
}

func assertNotBrand(t *testing.T, brand string, number string) {
	assert.NotEqual(t, brand, DetectBrand(number), "%s Should not detect %s for %s", caller(), brand, number)
}

func TestDetectDankort(t *testing.T) {
	assertBrand(t, "dankort", "5019717010103742")
}

func TestDetectVisaDankortAsVisa(t *testing.T) {
	assertBrand(t, "visa", "4571100000000000")
}

func TestDetectElectronDkAsVisa(t *testing.T) {
	assertBrand(t, "visa", "4175001000000000")
}

func TestDetectDinersClub(t *testing.T) {
	assertBrand(t, "diners_club", "36148010000000")
}

func TestDetectDinersClubDk(t *testing.T) {
	assertBrand(t, "diners_club", "30401000000000")
}

func TestDetectMaestroDkAsMaestro(t *testing.T) {
	assertBrand(t, "maestro", "6769271000000000")
}

func maestroCardNumbers() []string {
	return []string{
		"5000000000000000", "5099999999999999", "5600000000000000",
		"5899999999999999", "6000000000000000", "6999999999999999",
		"6761999999999999", "6763000000000000", "5038999999999999",
	}
}

func nonMaestroCardNumbers() []string {
	return []string{
		"4999999999999999", "5100000000000000", "5599999999999999",
		"5900000000000000", "5999999999999999", "7000000000000000",
	}
}

func TestDetectMaestroCards(t *testing.T) {
	assert.Equal(t, "maestro", DetectBrand("5020100000000000"))

	for _, number := range maestroCardNumbers() {
		assertBrand(t, "maestro", number)
	}

	for _, number := range nonMaestroCardNumbers() {
		assertNotBrand(t, "maestro", number)
	}
}

func TestDetectMasterCard(t *testing.T) {
	assertBrand(t, "master", "6771890000000000")
	assertBrand(t, "master", "5413031000000000")
}

func TestDetectForbrugsforeningen(t *testing.T) {
	assertBrand(t, "forbrugsforeningen", "6007221000000000")
}

func TestDetectLaser(t *testing.T) {
	// 16 digits
	assertBrand(t, "laser", "6304985028090561")

	// 18 digits
	assertBrand(t, "laser", "630498502809056151")

	// 19 digits
	assertBrand(t, "laser", "6304985028090561515")

	// 15 digits (not a laser)
	assertNotBrand(t, "laser", "630498502809056")

	// 17 digits (not a laser)
	assertNotBrand(t, "laser", "63049850280905615")

	// Alternate format
	assertBrand(t, "laser", "6706950000000000000")

	// Alternate format (16 digits)
	assertBrand(t, "laser", "6706123456789012")

	// New format (16 digits)
	assertBrand(t, "laser", "6709123456789012")

	// Ulster bank (Irelang) with 12 digits
	assertBrand(t, "laser", "677117111234")
}

func TestDetectingFullRangeOfMaestroCardNumbers(t *testing.T) {
	maestro := "50000000000"

	assertNotBrand(t, "maestro", maestro)

	for len(maestro) < 19 {
		maestro += "0"
		assertBrand(t, "maestro", maestro)
	}

	maestro += "0"
	assertNotBrand(t, "maestro", maestro)
}

func TestDetectDiscover(t *testing.T) {
	assertBrand(t, "discover", "6011000000000000")
	assertBrand(t, "discover", "6500000000000000")
	assertBrand(t, "discover", "6221260000000000")
	assertBrand(t, "discover", "6450000000000000")

	assertNotBrand(t, "discover", "6010000000000000")
	assertNotBrand(t, "discover", "6600000000000000")
}

func Test16DigitMaestroUK(t *testing.T) {
	assertBrand(t, "switch", "6759000000000000")
}

func Test18DigitMaestroUK(t *testing.T) {
	assertBrand(t, "switch", "675900000000000000")
}

func Test19DigitMaestroUK(t *testing.T) {
	assertBrand(t, "switch", "6759000000000000000")
}
