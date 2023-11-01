package gmath

import (
	"math"
	"math/rand"
)

func _RandomInt64(minInclusive int64, maxExclusive int64) int64 {
	if minInclusive == maxExclusive {
		return minInclusive
	}

	if minInclusive > maxExclusive {
		minInclusive, maxExclusive = maxExclusive, minInclusive
	}
	return rand.Int63n(maxExclusive-minInclusive) + minInclusive
}

// RandomInt [minInclusive,maxExclusive) if min==max return min
func RandomInt(minInclusive int, maxExclusive int) int {
	return int(_RandomInt64(int64(minInclusive), int64(maxExclusive)))
}

// RandomInt32 [minInclusive,maxExclusive) if min==max return min
func RandomInt32(minInclusive int32, maxExclusive int32) int32 {
	return int32(_RandomInt64(int64(minInclusive), int64(maxExclusive)))
}

// RandomUInt32 [minInclusive,maxExclusive) if min==max return min
func RandomUInt32(minInclusive uint32, maxExclusive uint32) uint32 {
	return uint32(_RandomInt64(int64(minInclusive), int64(maxExclusive)))
}

// RandomFloat32 [0,1.0]
func RandomFloat32() float32 {
	var max int32 = math.MaxInt32
	result := rand.Int31n(max)
	max = max - 1

	if result == (max - 1) {
		return 1.0
	}

	v := float64(result) / float64(max)
	return float32(v)
}

// RandomFloat64 [0,1.0]
func RandomFloat64() float64 {
	var max int32 = math.MaxInt32
	result := rand.Int31n(max)
	max = max - 1

	if result == (max - 1) {
		return 1.0
	}

	return float64(result) / float64(max)
}

// RandomInsideUnitCircle https://zhuanlan.zhihu.com/p/447898464
func RandomInsideUnitCircle() Vector2 {
	return RandomInsideCircle(0, 1)
}

// RandomInsideUnitCircle https://zhuanlan.zhihu.com/p/447898464
func RandomInsideCircle(minRadius, maxRadius float32) Vector2 {
	random_theta := rand.Float64() * PI2 //[0,1.0)

	//生成距离圆心的长度
	random_r := math.Sqrt(RandomFloat64()) //[0,1.0]
	random_r = float64(maxRadius-minRadius)*random_r + float64(minRadius)

	x := math.Cos(random_theta) * random_r
	y := math.Sin(random_theta) * random_r
	return Vector2{float32(x), float32(y)}
}
