package omh

func Omh(v, r, i float64) float64 {
	if v == 0 && r > 0 && i > 0 {
		return i * r
	} else if v > 0 && r == 0 && i > 0 {
		return v / i
	} else if v > 0 && r > 0 && i == 0 {
		return v / r
	} else {
		return 0
	}
}
