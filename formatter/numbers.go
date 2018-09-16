package formatter

import "math"

// Round2F rounds decimal number to two decimals
func Round2F(n float64) float64 {
	return math.Round(n*100) / 100
}
