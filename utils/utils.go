package utils

const HectogramsToKgConversionRate = float32(10)

// Converts hectograms to kilograms using a conversion rate of 10.
// The conversion rate represents how many kilograms correspond to one hectogram.
func HectogramsToKg(initial float32) float32 {
	return initial / HectogramsToKgConversionRate
}
