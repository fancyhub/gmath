package gmath

import "math"

const (
	Epsilon = math.SmallestNonzeroFloat32
)

func F32IsNaN(v float32) bool {
	return math.IsNaN(float64(v))
}

func F32Equal(v1, v2 float32) bool {
	return F32IsZero(v1 - v2)
}

func F32IsZero(v float32) bool {
	return F32Abs(v) < Epsilon
}

func F32Abs(v float32) float32 {
	switch {
	case v < 0:
		return -v
	case v == 0:
		return 0 // return correctly abs(-0)
	}
	return v
}

func F32Sqrt(v float32) float32 {
	return float32(math.Sqrt(float64(v)))
}

func F32Clamp(v, min, max float32) float32 {
	if v < min {
		return min
	} else if v > max {
		return max
	} else {
		return v
	}
}

func F32MoveTowards(from float32, to float32, maxDelta float32) float32 {
	if F32Abs(to-from) <= maxDelta {
		return to
	}
	if to > from {
		return from + maxDelta
	} else {
		return from - maxDelta
	}
}
