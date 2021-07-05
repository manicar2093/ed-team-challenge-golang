package ifloats

import "math"

// Floor2Positions32 floors the given float value to two 0 positions
func Floor2Positions32(v float64) float32 {
	return float32(math.Floor(v*100) / 100)
}
