package gmath

type Matrix4 struct {
	m00 float32
	m10 float32
	m20 float32
	m30 float32

	m01 float32
	m11 float32
	m21 float32
	m31 float32

	m02 float32
	m12 float32
	m22 float32
	m32 float32

	m03 float32
	m13 float32
	m23 float32
	m33 float32
}

func Matrix4Identity() Matrix4 {
	result := Matrix4{}
	result.m00 = 1
	result.m01 = 0
	result.m02 = 0
	result.m03 = 0

	result.m10 = 0
	result.m11 = 1
	result.m12 = 0
	result.m13 = 0

	result.m20 = 0
	result.m21 = 0
	result.m22 = 1
	result.m23 = 0

	result.m30 = 0
	result.m31 = 0
	result.m32 = 0
	result.m33 = 1

	return result
}

func Matrix4TRS(pos Vector3, rot Quaternion, scale Vector3) Matrix4 {
	this := Matrix4FromRotation(rot)

	this.m00 *= scale.X
	this.m10 *= scale.X
	this.m20 *= scale.X

	this.m01 *= scale.Y
	this.m11 *= scale.Y
	this.m21 *= scale.Y

	this.m02 *= scale.Z
	this.m12 *= scale.Z
	this.m22 *= scale.Z

	this.m03 = pos.X
	this.m13 = pos.Y
	this.m23 = pos.Z

	return this
}

func Matrix4FromRotation(q Quaternion) Matrix4 {
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

	result := Matrix4{}
	result.m00 = 1 - (num5 + num6)
	result.m10 = num7 + num12
	result.m20 = num8 - num11
	result.m30 = 0
	result.m01 = num7 - num12
	result.m11 = 1 - (num4 + num6)
	result.m21 = num9 + num10
	result.m31 = 0
	result.m02 = num8 + num11
	result.m12 = num9 - num10
	result.m22 = 1 - (num4 + num5)
	result.m32 = 0
	result.m03 = 0
	result.m13 = 0
	result.m23 = 0
	result.m33 = 1
	return result
}

func Matrix4FromTranslate(vector Vector3) Matrix4 {
	result := Matrix4{}
	result.m00 = 1
	result.m01 = 0
	result.m02 = 0
	result.m03 = vector.X

	result.m10 = 0
	result.m11 = 1
	result.m12 = 0
	result.m13 = vector.Y

	result.m20 = 0
	result.m21 = 0
	result.m22 = 1
	result.m23 = vector.Z

	result.m30 = 0
	result.m31 = 0
	result.m32 = 0
	result.m33 = 1
	return result
}

func Matrix4FromScale(vector Vector3) Matrix4 {
	result := Matrix4{}
	result.m00 = vector.X
	result.m01 = 0
	result.m02 = 0
	result.m03 = 0
	result.m10 = 0
	result.m11 = vector.Y
	result.m12 = 0
	result.m13 = 0
	result.m20 = 0
	result.m21 = 0
	result.m22 = vector.Z
	result.m23 = 0
	result.m30 = 0
	result.m31 = 0
	result.m32 = 0
	result.m33 = 1
	return result
}

func (matrix *Matrix4) GetColumn(index int) Vector4 {
	switch index {
	case 0:
		return Vector4{matrix.m00, matrix.m10, matrix.m20, matrix.m30}
	case 1:
		return Vector4{matrix.m01, matrix.m11, matrix.m21, matrix.m31}
	case 2:
		return Vector4{matrix.m02, matrix.m12, matrix.m22, matrix.m32}
	case 3:
		return Vector4{matrix.m03, matrix.m13, matrix.m23, matrix.m33}
	default:
		panic("index out of range")
	}
}
func (matrix *Matrix4) GetRow(index int) Vector4 {
	switch index {
	case 0:
		return Vector4{matrix.m00, matrix.m01, matrix.m02, matrix.m03}
	case 1:
		return Vector4{matrix.m10, matrix.m11, matrix.m12, matrix.m13}
	case 2:
		return Vector4{matrix.m20, matrix.m21, matrix.m22, matrix.m23}
	case 3:
		return Vector4{matrix.m30, matrix.m31, matrix.m32, matrix.m33}
	default:
		panic("index out of range")
	}
}
func (matrix *Matrix4) GetPosition() Vector3 {
	return Vector3{matrix.m03, matrix.m13, matrix.m23}
}

func (left *Matrix4) Multiply(right *Matrix4) Matrix4 {
	result := Matrix4{}
	result.m00 = left.m00*right.m00 + left.m01*right.m10 + left.m02*right.m20 + left.m03*right.m30
	result.m01 = left.m00*right.m01 + left.m01*right.m11 + left.m02*right.m21 + left.m03*right.m31
	result.m02 = left.m00*right.m02 + left.m01*right.m12 + left.m02*right.m22 + left.m03*right.m32
	result.m03 = left.m00*right.m03 + left.m01*right.m13 + left.m02*right.m23 + left.m03*right.m33
	result.m10 = left.m10*right.m00 + left.m11*right.m10 + left.m12*right.m20 + left.m13*right.m30
	result.m11 = left.m10*right.m01 + left.m11*right.m11 + left.m12*right.m21 + left.m13*right.m31
	result.m12 = left.m10*right.m02 + left.m11*right.m12 + left.m12*right.m22 + left.m13*right.m32
	result.m13 = left.m10*right.m03 + left.m11*right.m13 + left.m12*right.m23 + left.m13*right.m33
	result.m20 = left.m20*right.m00 + left.m21*right.m10 + left.m22*right.m20 + left.m23*right.m30
	result.m21 = left.m20*right.m01 + left.m21*right.m11 + left.m22*right.m21 + left.m23*right.m31
	result.m22 = left.m20*right.m02 + left.m21*right.m12 + left.m22*right.m22 + left.m23*right.m32
	result.m23 = left.m20*right.m03 + left.m21*right.m13 + left.m22*right.m23 + left.m23*right.m33
	result.m30 = left.m30*right.m00 + left.m31*right.m10 + left.m32*right.m20 + left.m33*right.m30
	result.m31 = left.m30*right.m01 + left.m31*right.m11 + left.m32*right.m21 + left.m33*right.m31
	result.m32 = left.m30*right.m02 + left.m31*right.m12 + left.m32*right.m22 + left.m33*right.m32
	result.m33 = left.m30*right.m03 + left.m31*right.m13 + left.m32*right.m23 + left.m33*right.m33
	return result
}

func (matrix *Matrix4) Transpose() Matrix4 {
	result := Matrix4{}
	result.m00 = matrix.m00
	result.m01 = matrix.m10
	result.m02 = matrix.m20
	result.m03 = matrix.m30

	result.m10 = matrix.m01
	result.m11 = matrix.m11
	result.m12 = matrix.m21
	result.m13 = matrix.m31

	result.m20 = matrix.m02
	result.m21 = matrix.m12
	result.m22 = matrix.m22
	result.m23 = matrix.m32

	result.m30 = matrix.m03
	result.m31 = matrix.m13
	result.m32 = matrix.m23
	result.m33 = matrix.m33
	return result
}

// Inverse
func (matrix *Matrix4) Inverse() Matrix4 {
	inv := Matrix4{}
	inv.m00 = matrix.m11*matrix.m22*matrix.m33 -
		matrix.m11*matrix.m32*matrix.m23 -
		matrix.m12*matrix.m21*matrix.m33 +
		matrix.m12*matrix.m31*matrix.m23 +
		matrix.m13*matrix.m21*matrix.m32 -
		matrix.m13*matrix.m31*matrix.m22

	inv.m01 = -matrix.m01*matrix.m22*matrix.m33 +
		matrix.m01*matrix.m32*matrix.m23 +
		matrix.m02*matrix.m21*matrix.m33 -
		matrix.m02*matrix.m31*matrix.m23 -
		matrix.m03*matrix.m21*matrix.m32 +
		matrix.m03*matrix.m31*matrix.m22

	inv.m02 = matrix.m01*matrix.m12*matrix.m33 -
		matrix.m01*matrix.m32*matrix.m13 -
		matrix.m02*matrix.m11*matrix.m33 +
		matrix.m02*matrix.m31*matrix.m13 +
		matrix.m03*matrix.m11*matrix.m32 -
		matrix.m03*matrix.m31*matrix.m12

	inv.m03 = -matrix.m01*matrix.m12*matrix.m23 +
		matrix.m01*matrix.m22*matrix.m13 +
		matrix.m02*matrix.m11*matrix.m23 -
		matrix.m02*matrix.m21*matrix.m13 -
		matrix.m03*matrix.m11*matrix.m22 +
		matrix.m03*matrix.m21*matrix.m12

	inv.m10 = -matrix.m10*matrix.m22*matrix.m33 +
		matrix.m10*matrix.m32*matrix.m23 +
		matrix.m12*matrix.m20*matrix.m33 -
		matrix.m12*matrix.m30*matrix.m23 -
		matrix.m13*matrix.m20*matrix.m32 +
		matrix.m13*matrix.m30*matrix.m22

	inv.m11 = matrix.m00*matrix.m22*matrix.m33 -
		matrix.m00*matrix.m32*matrix.m23 -
		matrix.m02*matrix.m20*matrix.m33 +
		matrix.m02*matrix.m30*matrix.m23 +
		matrix.m03*matrix.m20*matrix.m32 -
		matrix.m03*matrix.m30*matrix.m22

	inv.m12 = -matrix.m00*matrix.m12*matrix.m33 +
		matrix.m00*matrix.m32*matrix.m13 +
		matrix.m02*matrix.m10*matrix.m33 -
		matrix.m02*matrix.m30*matrix.m13 -
		matrix.m03*matrix.m10*matrix.m32 +
		matrix.m03*matrix.m30*matrix.m12

	inv.m13 = matrix.m00*matrix.m12*matrix.m23 -
		matrix.m00*matrix.m22*matrix.m13 -
		matrix.m02*matrix.m10*matrix.m23 +
		matrix.m02*matrix.m20*matrix.m13 +
		matrix.m03*matrix.m10*matrix.m22 -
		matrix.m03*matrix.m20*matrix.m12

	inv.m20 = matrix.m10*matrix.m21*matrix.m33 -
		matrix.m10*matrix.m31*matrix.m23 -
		matrix.m11*matrix.m20*matrix.m33 +
		matrix.m11*matrix.m30*matrix.m23 +
		matrix.m13*matrix.m20*matrix.m31 -
		matrix.m13*matrix.m30*matrix.m21

	inv.m21 = -matrix.m00*matrix.m21*matrix.m33 +
		matrix.m00*matrix.m31*matrix.m23 +
		matrix.m01*matrix.m20*matrix.m33 -
		matrix.m01*matrix.m30*matrix.m23 -
		matrix.m03*matrix.m20*matrix.m31 +
		matrix.m03*matrix.m30*matrix.m21

	inv.m22 = matrix.m00*matrix.m11*matrix.m33 -
		matrix.m00*matrix.m31*matrix.m13 -
		matrix.m01*matrix.m10*matrix.m33 +
		matrix.m01*matrix.m30*matrix.m13 +
		matrix.m03*matrix.m10*matrix.m31 -
		matrix.m03*matrix.m30*matrix.m11

	inv.m23 = -matrix.m00*matrix.m11*matrix.m23 +
		matrix.m00*matrix.m21*matrix.m13 +
		matrix.m01*matrix.m10*matrix.m23 -
		matrix.m01*matrix.m20*matrix.m13 -
		matrix.m03*matrix.m10*matrix.m21 +
		matrix.m03*matrix.m20*matrix.m11

	inv.m30 = -matrix.m10*matrix.m21*matrix.m32 +
		matrix.m10*matrix.m31*matrix.m22 +
		matrix.m11*matrix.m20*matrix.m32 -
		matrix.m11*matrix.m30*matrix.m22 -
		matrix.m12*matrix.m20*matrix.m31 +
		matrix.m12*matrix.m30*matrix.m21

	inv.m31 = matrix.m00*matrix.m21*matrix.m32 -
		matrix.m00*matrix.m31*matrix.m22 -
		matrix.m01*matrix.m20*matrix.m32 +
		matrix.m01*matrix.m30*matrix.m22 +
		matrix.m02*matrix.m20*matrix.m31 -
		matrix.m02*matrix.m30*matrix.m21

	inv.m32 = -matrix.m00*matrix.m11*matrix.m32 +
		matrix.m00*matrix.m31*matrix.m12 +
		matrix.m01*matrix.m10*matrix.m32 -
		matrix.m01*matrix.m30*matrix.m12 -
		matrix.m02*matrix.m10*matrix.m31 +
		matrix.m02*matrix.m30*matrix.m11

	inv.m33 = matrix.m00*matrix.m11*matrix.m22 -
		matrix.m00*matrix.m21*matrix.m12 -
		matrix.m01*matrix.m10*matrix.m22 +
		matrix.m01*matrix.m20*matrix.m12 +
		matrix.m02*matrix.m10*matrix.m21 -
		matrix.m02*matrix.m20*matrix.m11

	det := matrix.m00*inv.m00 + matrix.m10*inv.m01 + matrix.m20*inv.m02 + matrix.m30*inv.m03

	if det == 0 {
		//panic
		return Matrix4Identity()
	}

	detInverse := 1.0 / det
	inv.m00 *= detInverse
	inv.m01 *= detInverse
	inv.m02 *= detInverse
	inv.m03 *= detInverse

	inv.m10 *= detInverse
	inv.m11 *= detInverse
	inv.m12 *= detInverse
	inv.m13 *= detInverse

	inv.m20 *= detInverse
	inv.m21 *= detInverse
	inv.m22 *= detInverse
	inv.m23 *= detInverse

	inv.m30 *= detInverse
	inv.m31 *= detInverse
	inv.m32 *= detInverse
	inv.m33 *= detInverse
	return inv
}

// MultiplyPoint3 slow generic
func (matrix *Matrix4) MultiplyPoint3(point Vector3) Vector3 {
	result := Vector3{}
	result.X = matrix.m00*point.X + matrix.m01*point.Y + matrix.m02*point.Z + matrix.m03
	result.Y = matrix.m10*point.X + matrix.m11*point.Y + matrix.m12*point.Z + matrix.m13
	result.Z = matrix.m20*point.X + matrix.m21*point.Y + matrix.m22*point.Z + matrix.m23
	num := matrix.m30*point.X + matrix.m31*point.Y + matrix.m32*point.Z + matrix.m33
	num = 1 / num
	result.X *= num
	result.Y *= num
	result.Z *= num
	return result
}

// MultiplyPoint3x4 fast
func (matrix *Matrix4) MultiplyPoint3x4(point Vector3) Vector3 {
	result := Vector3{}
	result.X = matrix.m00*point.X + matrix.m01*point.Y + matrix.m02*point.Z + matrix.m03
	result.Y = matrix.m10*point.X + matrix.m11*point.Y + matrix.m12*point.Z + matrix.m13
	result.Z = matrix.m20*point.X + matrix.m21*point.Y + matrix.m22*point.Z + matrix.m23
	return result
}

func (matrix *Matrix4) MultiplyPoint4(vector Vector4) Vector4 {
	result := Vector4{}
	result.X = matrix.m00*vector.X + matrix.m01*vector.Y + matrix.m02*vector.Z + matrix.m03*vector.W
	result.Y = matrix.m10*vector.X + matrix.m11*vector.Y + matrix.m12*vector.Z + matrix.m13*vector.W
	result.Z = matrix.m20*vector.X + matrix.m21*vector.Y + matrix.m22*vector.Z + matrix.m23*vector.W
	result.W = matrix.m30*vector.X + matrix.m31*vector.Y + matrix.m32*vector.Z + matrix.m33*vector.W
	return result
}

func (matrix *Matrix4) MultiplyDir3(dir Vector3) Vector3 {
	result := Vector3{}
	result.X = matrix.m00*dir.X + matrix.m01*dir.Y + matrix.m02*dir.Z
	result.Y = matrix.m10*dir.X + matrix.m11*dir.Y + matrix.m12*dir.Z
	result.Z = matrix.m20*dir.X + matrix.m21*dir.Y + matrix.m22*dir.Z
	return result
}
