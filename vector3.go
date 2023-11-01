package gmath

import (
	"math"
)

type Vector3 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func V3Zero() Vector3    { return Vector3{0, 0, 0} }
func V3One() Vector3     { return Vector3{1, 1, 1} }
func V3Up() Vector3      { return Vector3{0, 1, 0} }
func V3Down() Vector3    { return Vector3{0, -1, 0} }
func V3Left() Vector3    { return Vector3{-1, 0, 0} }
func V3Right() Vector3   { return Vector3{1, 0, 0} }
func V3Forward() Vector3 { return Vector3{0, 0, 1} }
func V3Back() Vector3    { return Vector3{0, 0, -1} }

func (v3 *Vector3) Set(x, y, z float32) {
	v3.X = x
	v3.Y = y
	v3.Z = z
}

func (v3 Vector3) XZ() Vector2 {
	return Vector2{v3.X, v3.Z}
}
func (v3 Vector3) X0Z() Vector3 {
	return Vector3{v3.X, 0, v3.Z}
}

func (v3 *Vector3) Add(v Vector3) Vector3 {
	return Vector3{v3.X + v.X, v3.Y + v.Y, v3.Z + v.Z}
}

func (v3 *Vector3) AddSelf(v Vector3) {
	v3.X += v.X
	v3.Y += v.Y
	v3.Z += v.Z
}

func (v3 Vector3) Substract(v Vector3) Vector3 {
	return Vector3{v3.X - v.X, v3.Y - v.Y, v3.Z - v.Z}
}

func (v3 Vector3) Scale(v float32) Vector3 {
	return Vector3{v3.X * v, v3.Y * v, v3.Z * v}
}

func (v3 *Vector3) ScaleSelf(v float32) {
	v3.X *= v
	v3.Y *= v
	v3.Z *= v
}

func (v3 Vector3) ScaleV3(v Vector3) Vector3 {
	return Vector3{v3.X * v.X, v3.Y * v.X, v3.Z * v.X}
}

func (v3 Vector3) Dot(v Vector3) float32 {
	return v3.X*v.X + v3.Y*v.Y + v3.Z*v.Z
}

func (v3 Vector3) Cross(v Vector3) Vector3 {
	temp := Vector3{}
	temp.X = v3.Y*v.Z - v3.Z*v.Y
	temp.Y = v3.Z*v.X - v3.X*v.Z
	temp.Z = v3.X*v.Y - v3.Y*v.X
	return temp
}

func (v3 Vector3) Magnitude() float32 {
	return F32Sqrt(v3.X*v3.X + v3.Y*v3.Y + v3.Z*v3.Z)
}

func (v3 Vector3) SqrMagnitude() float32 {
	return v3.X*v3.X + v3.Y*v3.Y + v3.Z*v3.Z
}

func (v3 Vector3) Normalize() Vector3 {
	v3.NormalizeSelf()
	return v3
}

func (v3 *Vector3) NormalizeSelf() float32 {
	var magn = v3.Magnitude()
	if F32IsZero(magn) {
		return 0
	}
	v3.X = v3.X / magn
	v3.Y = v3.Y / magn
	v3.Z = v3.Z / magn
	return magn
}

func (v3 Vector3) IsZero() bool {
	return F32IsZero(v3.SqrMagnitude())
}

func (v3 Vector3) Equal(v Vector3) bool {
	return F32Equal(v3.X, v.X) && F32Equal(v3.Y, v.Y) && F32Equal(v3.Z, v.Z)
}

func (v3 Vector3) IsValid() bool {
	return !(math.IsNaN(float64(v3.X)) || math.IsNaN(float64(v3.Y)) || math.IsNaN(float64(v3.Z)))
}

//========================

func V3Lerp(from Vector3, to Vector3, t float32) Vector3 {
	t = F32Clamp01(t)
	return Vector3{from.X + (to.X-from.X)*t, from.Y + (to.Y-from.Y)*t, from.Z + (to.Z-from.Z)*t}
}

func V3LerpUnclamped(from Vector3, to Vector3, t float32) Vector3 {
	return Vector3{from.X + (to.X-from.X)*t, from.Y + (to.Y-from.Y)*t, from.Z + (to.Z-from.Z)*t}
}

func V3MoveTowards(current Vector3, target Vector3, maxDistanceDelta float32) Vector3 {
	num := target.X - current.X
	num2 := target.Y - current.Y
	num3 := target.Z - current.Z
	num4 := num*num + num2*num2 + num3*num3

	if num4 == 0 || (maxDistanceDelta >= 0 && num4 <= maxDistanceDelta*maxDistanceDelta) {
		return target
	}
	num5 := F32Sqrt(num4)
	return Vector3{current.X + num/num5*maxDistanceDelta, current.Y + num2/num5*maxDistanceDelta, current.Z + num3/num5*maxDistanceDelta}
}

// V3Angle [0,180]
func V3Angle(from Vector3, to Vector3) AngleDegree {
	num := from.SqrMagnitude() * to.SqrMagnitude()
	if num < 1e-7 {
		return 0
	}
	num = F32Sqrt(num)

	num2 := F32Clamp(from.Dot(to)/num, -1, 1)
	return Acos(num2).ToDegrees()
}

// V3SignedAngle [-180,180]
func V3SignedAngle(from, to, axis Vector3) AngleDegree {
	num := V3Angle(from, to)
	num2 := from.Y*to.Z - from.Z*to.Y
	num3 := from.Z*to.X - from.X*to.Z
	num4 := from.X*to.Y - from.Y*to.X
	num5 := F32Sign(axis.X*num2 + axis.Y*num3 + axis.Z*num4)
	return num * AngleDegree(num5)
}

func V3SignedAngleY(from, to Vector3) AngleDegree {
	return V3SignedAngle(from, to, V3Up())
}

func V3Distance(from, to Vector3) float32 {
	return to.Substract(from).Magnitude()
}

func V3DistanceSqr(from, to Vector3) float32 {
	return to.Substract(from).SqrMagnitude()
}

func V3DistanceXZ(from, to Vector3) float32 {
	return to.Substract(from).XZ().Magnitude()
}

func V3DistanceXZSqr(from, to Vector3) float32 {
	return to.Substract(from).XZ().SqrMagnitude()
}
