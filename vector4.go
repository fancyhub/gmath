package gmath

type Vector4 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
}

func (v4 Vector4) XYZ() Vector3 {
	return Vector3{v4.X, v4.Y, v4.Z}
}
