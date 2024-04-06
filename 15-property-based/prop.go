package prop

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumericals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result = strings.Builder{}
	for _, numerical := range allRomanNumericals {
		for arabic >= numerical.Value {
			result.WriteString(numerical.Symbol)
			arabic -= numerical.Value
		}
	}
	return result.String()
}

func convertToArabic(roman string) int {
	arabic := 0
	for _, numerical := range allRomanNumericals {
		for strings.HasPrefix(roman, numerical.Symbol) {
			arabic += numerical.Value
			roman = strings.TrimPrefix(roman, numerical.Symbol)
		}
	}
	return arabic
}
