package gmath

import "math"

type Vector2 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func V2Zero() Vector2 { return Vector2{0, 0} }
func V2One() Vector2  { return Vector2{1, 1} }

func (v2 *Vector2) Set(x, y float32) {
	v2.X = x
	v2.Y = y
}

func (v2 *Vector2) Dot(v Vector2) float32 {
	return v2.X*v.X + v2.Y*v.Y
}

func (v2 *Vector2) Add(v Vector2) Vector2 {
	return Vector2{v2.X + v.X, v2.Y + v.Y}
}

func (v2 *Vector2) Substract(v Vector2) Vector2 {
	return Vector2{v2.X - v.X, v2.Y - v.Y}
}

func (v2 *Vector2) Scale(v float32) Vector2 {
	return Vector2{v2.X * v, v2.Y * v}
}

func (v2 *Vector2) ScaleSelf(v float32) {
	v2.X *= v
	v2.Y *= v
}

func (v2 *Vector2) ScaleV2(v Vector2) Vector2 {
	return Vector2{v2.X * v.X, v2.Y * v.Y}
}

func (v2 *Vector2) Magnitude() float32 {
	return F32Sqrt(v2.X*v2.X + v2.Y*v2.Y)
}

func (v2 *Vector2) SqrMagnitude() float32 {
	return v2.X*v2.X + v2.Y*v2.Y
}

func (v2 *Vector2) Normalize() Vector2 {
	temp := *v2
	temp.NormalizeSelf()
	return temp
}

func (v2 *Vector2) NormalizeSelf() float32 {
	var magn = v2.Magnitude()
	if F32IsZero(magn) {
		return 0
	}
	v2.X = v2.X / magn
	v2.Y = v2.Y / magn
	return magn
}

func (v2 *Vector2) IsZero() bool {
	return F32IsZero(v2.SqrMagnitude())
}

func (v2 *Vector2) Equal(v Vector3) bool {
	return F32Equal(v2.X, v.X) && F32Equal(v2.Y, v.Y)
}

func (v2 *Vector2) IsValid() bool {
	return !(math.IsNaN(float64(v2.X)) || math.IsNaN(float64(v2.Y)))
}
