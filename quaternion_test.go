package gmath

import "testing"

func TestQuaternion(t *testing.T) {

	eulerAngle := Vector3{10, 20, 30}
	eulerQuaternion := QuaternionFromEulerAngle(eulerAngle)
	if !eulerQuaternion.Equal(Quaternion{0.12767944, 0.14487813, 0.23929833, 0.9515485}) {
		t.Error("QuaternionFromEulerAngle")
	}

	eulerAngle2 := eulerQuaternion.ToEulerAngle()
	if !eulerAngle.Equal(eulerAngle2) {
		t.Error("ToEulerAngle")
	}

	eulerQuaternion2 := QuaternionFromEulerAngle(eulerAngle2)
	if !eulerQuaternion.Equal(eulerQuaternion2) {
		t.Error("QuaternionFromEulerAngle 2")
	}

	degree := QuaternionAngle(eulerQuaternion, QuaternionIdentity())
	if degree != 35.817104 {
		t.Error("QuaternionAngle")
	}
}
