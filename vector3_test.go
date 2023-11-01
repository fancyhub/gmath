package gmath

import "testing"

func TestVector3(t *testing.T) {
	a := V3Angle(V3Forward(), V3Back())
	if a != 180 {
		t.Error("V3Angle")
	}

	a = V3Angle(V3Forward(), V3Forward())
	if a != 0 {
		t.Error("V3Angle")
	}

	a = V3SignedAngle(V3Forward(), Vector3{3, 0, 3}.Normalize(), V3Up())
	if a != 45 {
		t.Error("V3SignedAngle")
	}

	a = V3SignedAngle(V3Forward(), Vector3{-3, 0, 3}.Normalize(), V3Up())
	if a != -45 {
		t.Error("V3SignedAngle")
	}
}
