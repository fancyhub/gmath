package gmath

import (
	"math"
)

const (
	_QuaternionEpsilon = 0.001
)

type Quaternion struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
}

func QuaternionIdentity() Quaternion { return Quaternion{0, 0, 0, 1} }

func (left Quaternion) Multiply(right Quaternion) Quaternion {
	x := left.W*right.X + left.X*right.W + left.Y*right.Z - left.Z*right.Y
	y := left.W*right.Y + left.Y*right.W + left.Z*right.X - left.X*right.Z
	z := left.W*right.Z + left.Z*right.W + left.X*right.Y - left.Y*right.X
	w := left.W*right.W - left.X*right.X - left.Y*right.Y - left.Z*right.Z
	return Quaternion{x, y, z, w}
}

func (q Quaternion) Inverse() Quaternion {
	q.InverseSelf()
	return q
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{-q.X, -q.Y, -q.Z, q.W}
}

func (q *Quaternion) NormalizeSelf() bool {
	n := q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W

	if F32Equal2(n, 1.0, _QuaternionEpsilon) {
		return true
	}
	if n < _QuaternionEpsilon { //Too small
		return false
	}
	n = 1.0 / F32Sqrt(n)
	q.X *= n
	q.Y *= n
	q.Z *= n
	q.W *= n
	return true
}

func (q *Quaternion) InverseSelf() bool {
	q.X, q.Y, q.Z = -q.X, -q.Y, -q.Z
	return q.NormalizeSelf()
}

func (q *Quaternion) MultiplyV3(v Vector3) Vector3 {
	num := q.X * 2
	num2 := q.Y * 2
	num3 := q.Z * 2
	num4 := q.X * num
	num5 := q.Y * num2
	num6 := q.Z * num3
	num7 := q.X * num2
	num8 := q.X * num3
	num9 := q.Y * num3
	num10 := q.W * num
	num11 := q.W * num2
	num12 := q.W * num3
	result := Vector3{}
	result.X = (1-(num5+num6))*v.X + (num7-num12)*v.Y + (num8+num11)*v.Z
	result.Y = (num7+num12)*v.X + (1-(num4+num6))*v.Y + (num9-num10)*v.Z
	result.Z = (num8-num11)*v.X + (num9+num10)*v.Y + (1-(num4+num5))*v.Z
	return result
}

func (a *Quaternion) Dot(b Quaternion) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func (q *Quaternion) Equal(q2 Quaternion) bool {
	d := q.Dot(q2)
	return F32Equal2(d, 1.0, _QuaternionEpsilon)
}

func (q *Quaternion) IsZero() bool {
	return F32IsZero(q.X) && F32IsZero(q.Y) && F32IsZero(q.Z) && F32IsZero(q.W)
}

// ToEulerAngle YXZ, degree
func (q *Quaternion) ToEulerAngle() Vector3 {
	var eulerX, eulerY, eulerZ AngleRadian
	xx := 2 * q.X * q.X
	yy := 2 * q.Y * q.Y
	zz := 2 * q.Z * q.Z
	xy := 2 * q.X * q.Y
	xz := 2 * q.X * q.Z
	yz := 2 * q.Y * q.Z
	wx := 2 * q.W * q.X
	wy := 2 * q.W * q.Y
	wz := 2 * q.W * q.Z

	matrix3_0 := 1 - (yy + zz)
	matrix3_1 := xy + wz
	// matrix3_2 := xz - wy
	matrix3_3 := xy - wz
	matrix3_4 := 1 - (xx + zz)
	// matrix3_5 := yz + wx
	matrix3_6 := xz + wy
	matrix3_7 := yz - wx
	matrix3_8 := 1 - (xx + yy)

	threshold := yz - wx

	if threshold >= 0.999 {
		eulerX = -PI * 0.5
		eulerY = Atan2(-matrix3_3, matrix3_0)
		eulerZ = 0
	} else if threshold <= -0.999 {
		eulerX = PI * 0.5
		eulerY = Atan2(matrix3_3, matrix3_0)
		eulerZ = 0
	} else {
		eulerX = Asin(-matrix3_7)
		eulerY = Atan2(matrix3_6, matrix3_8)
		eulerZ = Atan2(matrix3_1, matrix3_4)
	}

	return Vector3{
		X: eulerX.ToDegrees().Normalize().ToFloat32(),
		Y: eulerY.ToDegrees().Normalize().ToFloat32(),
		Z: eulerZ.ToDegrees().Normalize().ToFloat32(),
	}
}

func QuaternionAngle(a, b Quaternion) AngleDegree {
	num := F32Min(F32Abs(a.Dot(b)), 1)
	if num > 0.999999 {
		return 0
	} else {
		return AngleDegree(Acos(num) * 2 * 57.29578)
	}
}

func QuaternionAngleAxis(angle AngleRadian, axis Vector3) Quaternion {
	ret := Quaternion{}

	halfAngle := 0.5 * angle
	sinValue := Sin(halfAngle)
	cosValue := Cos(halfAngle)

	axisN := axis.Normalize()
	ret.X = sinValue * axisN.X
	ret.Y = sinValue * axisN.Y
	ret.Z = sinValue * axisN.Z
	ret.W = cosValue

	return ret
}

func QuaternionLookRotation(forward Vector3, up Vector3) Quaternion {
	vector := forward.Normalize()
	vector2 := up.Cross(vector)
	vector2.NormalizeSelf()
	vector3 := vector.Cross(vector2)

	m00 := vector2.X
	m01 := vector2.Y
	m02 := vector2.Z
	m10 := vector3.X
	m11 := vector3.Y
	m12 := vector3.Z
	m20 := vector.X
	m21 := vector.Y
	m22 := vector.Z

	num8 := m00 + m11 + m22
	quaternion := Quaternion{}
	if num8 > 0 {
		num := float32(math.Sqrt(float64(num8 + 1.0)))
		quaternion.W = float32(num * 0.5)
		num = 0.5 / num
		quaternion.X = (m12 - m21) * num
		quaternion.Y = (m20 - m02) * num
		quaternion.Z = (m01 - m10) * num
		return quaternion
	}
	if m00 >= m11 && m00 >= m22 {
		num7 := float32(math.Sqrt(float64(1.0 + m00 - m11 - m22)))
		num4 := 0.5 / num7
		quaternion.X = 0.5 * num7
		quaternion.Y = (m01 + m10) * num4
		quaternion.Z = (m02 + m20) * num4
		quaternion.W = (m12 - m21) * num4
		return quaternion
	}
	if m11 > m22 {
		num6 := float32(math.Sqrt(float64(1.0 + m11 - m00 - m22)))
		num3 := 0.5 / num6
		quaternion.X = (m10 + m01) * num3
		quaternion.Y = 0.5 * num6
		quaternion.Z = (m21 + m12) * num3
		quaternion.W = (m20 - m02) * num3
		return quaternion
	}

	num5 := float32(math.Sqrt(float64(1.0 + m22 - m00 - m11)))
	num2 := 0.5 / num5
	quaternion.X = (m20 + m02) * num2
	quaternion.Y = (m21 + m12) * num2
	quaternion.Z = 0.5 * num5
	quaternion.W = (m01 - m10) * num2
	return quaternion

}

func QuaternionFromTo(fromVector, toVector Vector3) Quaternion {
	norm := math.Sqrt(float64(fromVector.SqrMagnitude() * toVector.SqrMagnitude()))
	cos_theta := float64(fromVector.Dot(toVector)) / norm
	half_cos := math.Sqrt(0.5 * (1 + cos_theta))

	w := fromVector.Cross(toVector)
	w = w.Scale(float32(1 / (norm * 2 * half_cos)))

	ret := Quaternion{}
	ret.X = w.X
	ret.Y = w.Y
	ret.Z = w.Z
	ret.W = float32(half_cos)
	return ret
}

func QuaternionLookFoward(forward Vector3) Quaternion {
	rightVector := V3Up().Cross(forward)
	upVector := forward.Cross(rightVector)
	return QuaternionLookRotation(forward, upVector)
}

// QuaternionFromEulerAngle YXZ, degree
func QuaternionFromEulerAngle(euler Vector3) Quaternion {
	qY := QuaternionAngleAxis(AngleRadian(euler.Y*Deg2Rad), V3Up())
	qX := QuaternionAngleAxis(AngleRadian(euler.X*Deg2Rad), V3Right())
	qZ := QuaternionAngleAxis(AngleRadian(euler.Z*Deg2Rad), V3Forward())

	return qY.Multiply(qX).Multiply(qZ)
}

func QuaternionLerp(from Quaternion, to Quaternion, t float32) Quaternion {
	t = F32Clamp01(t)
	return Quaternion{
		X: (1-t)*from.X + t*to.X,
		Y: (1-t)*from.Y + t*to.Y,
		Z: (1-t)*from.Z + t*to.Z,
		W: (1-t)*from.W + t*to.W,
	}
}

func QuaternionLerpUnclamped(from Quaternion, to Quaternion, t float32) Quaternion {
	return Quaternion{
		X: (1-t)*from.X + t*to.X,
		Y: (1-t)*from.Y + t*to.Y,
		Z: (1-t)*from.Z + t*to.Z,
		W: (1-t)*from.W + t*to.W,
	}
}
func QuaternionSlerp(from Quaternion, to Quaternion, t float32) Quaternion {
	return QuaternionSlerpUnclamped(from, to, F32Clamp01(t))
}

func QuaternionSlerpUnclamped(from Quaternion, to Quaternion, t float32) Quaternion {
	cos_theta := from.Dot(to)
	var sign float32 = 1.0
	if cos_theta < 0 {
		cos_theta = -cos_theta
		sign = -1
	}

	var c1, c2 float32
	if cos_theta > 1-_QuaternionEpsilon {
		c2 = t
		c1 = 1 - t
	} else {
		theta := Acos(cos_theta)
		sin_theta := Sin(theta)
		t_theta := theta.Multiply(t)
		inv_sin_theta := 1 / sin_theta
		c2 = Sin(t_theta) * inv_sin_theta
		c1 = Sin(theta-t_theta) * inv_sin_theta
	}

	c2 *= sign

	return Quaternion{
		X: from.X*c1 + to.X*c2,
		Y: from.Y*c1 + to.Y*c2,
		Z: from.Z*c1 + to.Z*c2,
		W: from.W*c1 + to.W*c2,
	}
}
