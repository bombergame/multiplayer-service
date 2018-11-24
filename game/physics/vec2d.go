package physics

type Vec2D struct {
	X Coordinate
	Y Coordinate
}

func GetVec2D(x, y Coordinate) Vec2D {
	return Vec2D{X: x, Y: y}
}

func GetVec2DZeros() Vec2D {
	return GetVec2D(0, 0)
}
