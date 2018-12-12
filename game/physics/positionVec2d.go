package physics

//go:generate easyjson

//easyjson:json
type PositionVec2D struct {
	Vec2D
}

func GetPositionVec2D(x, y Coordinate) PositionVec2D {
	return PositionVec2D{Vec2D: Vec2D{X: x, Y: y}}
}

func GetPositionVec2DZeros() PositionVec2D {
	return GetPositionVec2D(0, 0)
}

func (p PositionVec2D) Up(d Coordinate) PositionVec2D {
	return GetPositionVec2D(p.X, p.Y-d)
}

func (p PositionVec2D) Down(d Coordinate) PositionVec2D {
	return GetPositionVec2D(p.X, p.Y+d)
}

func (p PositionVec2D) Left(d Coordinate) PositionVec2D {
	return GetPositionVec2D(p.X-1, p.Y)
}

func (p PositionVec2D) Right(d Coordinate) PositionVec2D {
	return GetPositionVec2D(p.X+1, p.Y)
}
