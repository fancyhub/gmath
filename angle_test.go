package gmath

import "testing"

func TestAngle(t *testing.T) {
	v := Sin(30 * Deg2Rad)
	if v != 0.5 {
		t.Error("Sin")
	}

	v = Cos(60 * Deg2Rad)
	if !F32Equal(v, 0.5) {
		t.Error("Cos")
	}

	vDegree := AngleDegreeMoveTowards(40, -40, 60)
	if vDegree != -20 {
		t.Error("AngleDegreeMoveTowards")
	}

	vDegree = AngleDegreeMoveTowards(-40, 40, 60)
	if vDegree != 20 {
		t.Error("AngleDegreeMoveTowards")
	}
}
