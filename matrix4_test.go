package gmath

import (
	"testing"
)

func TestMatrix4(t *testing.T) {

	m1 := Matrix4TRS(Vector3{0, 1, 2}, QuaternionFromEulerAngle(Vector3{10, 20, 30}), V3One())
	m2 := m1.Inverse()
	m3 := m1.Multiply(&m2)
	v1 := Vector3{100, 200, 300}
	v2 := m3.MultiplyPoint3(v1)
	if v1.Substract(v2).Magnitude() > 0.1 {
		t.Error("Inverse & Multiply & MultiplyPoint3")
	}

	m4 := m1.Transpose()
	m5 := m4.Transpose()

	if m1 != m5 {
		t.Error("Transpose")
	}
}
