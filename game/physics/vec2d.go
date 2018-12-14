package physics

//go:generate easyjson

//easyjson:json
type Vec2D struct {
	X Coordinate `json:"x"`
	Y Coordinate `json:"y"`
}

func GetVec2D(x, y Coordinate) Vec2D {
	return Vec2D{X: x, Y: y}
}

func GetVec2DZeros() Vec2D {
	return GetVec2D(0, 0)
}
