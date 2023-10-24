package categorizebox

import "math"

func categorizeBox(length int, width int, height int, mass int) string {
	dimension := int(math.Pow10(4))
	isBulky := false
	isHeavy := false
	if length >= dimension || width >= dimension || height >= dimension || float64(length*width*height) >= math.Pow10(9) {
		isBulky = true
	}
	if mass >= 100 {
		isHeavy = true
	}
	if isBulky && isHeavy {
		return "Both"
	}
	if isBulky {
		return "Bulky"
	}
	if isHeavy {
		return "Heavy"
	}
	return "Neither"
}
