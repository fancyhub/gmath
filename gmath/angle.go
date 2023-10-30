package gmath

import "math"

const (
	PI      = math.Pi
	PI2     = math.Pi * 2
	Deg2Rad = math.Pi / 180
	Rad2Deg = 180 / math.Pi
)

type AngleDegree float32
type AngleRadian float32

func (angle AngleDegree) ToRadian() AngleRadian {
	return AngleRadian(float32(angle) * Deg2Rad)
}
func (angle AngleDegree) ToFloat32() float32 {
	return float32(angle)
}
func (angle AngleDegree) Multiply(v float32) AngleDegree {
	return AngleDegree(float32(angle) * v)
}

// Normalize [0,360]
func (angle AngleDegree) Normalize() AngleDegree {
	ret := float64(angle)
	ret -= math.Floor(ret/360.0) * 360.0
	if ret < 0 {
		ret = 0
	} else if ret > 360 {
		ret = 360
	}
	return AngleDegree(ret)
}

// NormalizeHalf [-180,180]
func (angle AngleDegree) NormalizeHalf() AngleDegree {
	ret := angle.Normalize()
	if ret > 180 {
		ret -= 360
	}
	return ret
}

func (angle AngleRadian) ToDegrees() AngleDegree {
	return AngleDegree(float32(angle) * Rad2Deg)
}
func (angle AngleRadian) ToFloat32() float32 {
	return float32(angle)
}

func (angle AngleRadian) Multiply(v float32) AngleRadian {
	return AngleRadian(float32(angle) * v)
}

// Normalize [0,TWO_PI]
func (angle AngleRadian) Normalize() AngleRadian {
	ret := float64(angle)
	ret -= math.Floor(ret/PI2) * PI2
	if ret < 0 {
		ret = 0
	} else if ret > PI2 {
		ret = PI2
	}
	return AngleRadian(ret)
}

// NormalizeHalf [-PI,PI]
func (angle AngleRadian) NormalizeHalf() AngleRadian {
	ret := angle.Normalize()
	if ret > PI {
		ret -= PI2
	}
	return ret
}

// AngleDegreeDelta  Result [-180,180]
func AngleDegreeDelta(from, to AngleDegree) AngleDegree {
	ret := to - from
	return ret.NormalizeHalf()
}

// AngleRadianDelta  Result [-PI,PI]
func AngleRadianDelta(from, to AngleRadian) AngleRadian {
	ret := to - from
	return ret.NormalizeHalf()
}

// AngleRadianMoveTowards return [0,TWO_PI]
func AngleRadianMoveTowards(from AngleRadian, to AngleRadian, maxDelta AngleRadian) AngleRadian {
	dtAngle := AngleRadianDelta(from, to)
	if -maxDelta < dtAngle && dtAngle < maxDelta {
		return to
	}
	to = from + dtAngle

	ret := AngleRadian(F32MoveTowards(from.ToFloat32(), to.ToFloat32(), maxDelta.ToFloat32()))
	return ret.Normalize()
}

// AngleDegreeMoveTowards return [0,TWO_PI]
func AngleDegreeMoveTowards(from AngleDegree, to AngleDegree, maxDelta AngleDegree) AngleDegree {
	dtAngle := AngleDegreeDelta(from, to)
	if -maxDelta < dtAngle && dtAngle < maxDelta {
		return to
	}
	to = from + dtAngle

	ret := AngleDegree(F32MoveTowards(from.ToFloat32(), to.ToFloat32(), maxDelta.ToFloat32()))
	return ret.Normalize()
}

func Sin(v AngleRadian) float32 {
	return float32(math.Sin(float64(v)))
}
func Asin(v float32) AngleRadian {
	return AngleRadian(math.Asin(float64(F32Clamp(v, -1, 1))))
}
func Cos(v AngleRadian) float32 {
	return float32(math.Cos(float64(v)))
}
func Acos(v float32) AngleRadian {
	return AngleRadian(math.Acos(float64(F32Clamp(v, -1, 1))))
}
func Tan(v AngleRadian) float32 {
	return float32(math.Tan(float64(v)))
}
func Atan(v float32) AngleRadian {
	return AngleRadian(math.Atan(float64(v)))
}
func Atan2(y, x float32) AngleRadian {
	return AngleRadian(math.Atan2(float64(y), float64(x)))
}
