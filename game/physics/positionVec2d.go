package physics

type PositionVec2D struct {
	Vec2D
}

func GetPositionVec2D(x, y Coordinate) PositionVec2D {
	return PositionVec2D{Vec2D: Vec2D{X: x, Y: y}}
}

func GetPositionVec2DZeros() PositionVec2D {
	return GetPositionVec2D(0, 0)
}

func (v *PositionVec2D) Translate(sp SpeedVec2D, t Time) *PositionVec2D {
	v.X += Coordinate(Float(sp.X) * Float(t))
	v.Y += Coordinate(Float(sp.Y) * Float(t))
	return v
}
