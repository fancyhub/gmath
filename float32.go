package gmath

import "math"

const (
	Epsilon = math.SmallestNonzeroFloat32
)

func F32IsNaN(v float32) bool {
	return math.IsNaN(float64(v))
}

func F32Equal(v1, v2 float32) bool {
	return F32Abs(v1-v2) < Epsilon
}

func F32Equal2(v1, v2, epsilon float32) bool {
	return F32Abs(v1-v2) < epsilon
}

func F32IsZero2(v, epsilon float32) bool {
	return F32Abs(v) < epsilon
}
func F32IsZero(v float32) bool {
	return F32Abs(v) < Epsilon
}

func F32Abs(v float32) float32 {
	if v < 0 {
		return -v
	} else if v == 0 {
		return 0
	} else {
		return v
	}
}
func F32Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
func F32Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
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

func F32Clamp01(v float32) float32 {
	if v < 0 {
		return 0
	} else if v > 1 {
		return 1
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

func F32LerpUnclamped(from float32, to float32, t float32) float32 {
	return (to-from)*t + from
}

func F32Lerp(from float32, to float32, t float32) float32 {
	return (to-from)*F32Clamp01(t) + from
}
