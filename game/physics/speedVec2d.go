package physics

type SpeedVec2D struct {
	Vec2D
}

func GetSpeedVec2D(x, y Speed) SpeedVec2D {
	return SpeedVec2D{Vec2D: Vec2D{X: Coordinate(x), Y: Coordinate(y)}}
}

func GetSpeedVec2DZeros() SpeedVec2D {
	return GetSpeedVec2D(0, 0)
}
