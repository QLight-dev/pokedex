package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const hectogramsToKgConversionRate = float32(10)

// Convert hectograms to kilograms.
func HectogramsToKg(initial float32) float32 {
	return initial / hectogramsToKgConversionRate
}

const decimetersToMetersConversionRate = float32(10)

// Convert decimeters to meters.
func DecimetersToMeters(initial float32) float32 {
	return initial / decimetersToMetersConversionRate
}

func FormatName(name string) string {
	// replace - to space to look more professional
	name = strings.ReplaceAll(name, "-", " ")

	// upercase each word to make more professional
	caser := cases.Title(language.English)
	formattedName := caser.String(name)
	return formattedName
}
