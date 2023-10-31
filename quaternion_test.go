package gmath

import "testing"

func TestQuaternion(t *testing.T) {

	eulerAngle := Vector3{10, 20, 30}
	eulerQuaternion := QuaternionFromEulerAngle(eulerAngle)
	if eulerQuaternion.Equal(Quaternion{0.12767944, 0.14487813, 0.23929833, 0.9515485}) {
		t.Log("QuaternionFromEulerAngle")
	} else {
		t.Error("QuaternionFromEulerAngle")
	}

	eulerAngle2 := eulerQuaternion.ToEulerAngle()
	if eulerAngle.Equal(eulerAngle2) {
		t.Log("ToEulerAngle")
	} else {
		t.Error("ToEulerAngle")
	}

	degree := QuaternionAngle(eulerQuaternion, QuaternionIdentity())
	t.Log(degree)
}
