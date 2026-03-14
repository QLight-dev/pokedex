package utils

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
