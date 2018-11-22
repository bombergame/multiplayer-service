package physics

type float float64

type Coordinate float
type Time float
type Speed float

type Vec2D struct {
	X Coordinate
	Y Coordinate
}

type PositionVec2D struct {
	Vec2D
}

type SpeedVec2D struct {
	Vec2D
}

func GetVec2D(x, y Coordinate) Vec2D {
	return Vec2D{X: x, Y: y}
}

func GetVec2DZeros() Vec2D {
	return GetVec2D(0, 0)
}

func GetPositionVec2D(x, y Coordinate) PositionVec2D {
	return PositionVec2D{Vec2D: Vec2D{X: x, Y: y}}
}

func GetPositionVec2DZeros() PositionVec2D {
	return GetPositionVec2D(0, 0)
}

func GetSpeedVec2D(x, y Speed) SpeedVec2D {
	return SpeedVec2D{Vec2D: Vec2D{X: Coordinate(x), Y: Coordinate(y)}}
}

func GetSpeedVec2DZeros() SpeedVec2D {
	return GetSpeedVec2D(0, 0)
}

func (v *PositionVec2D) Translate(sp SpeedVec2D, t Time) *PositionVec2D {
	v.X += Coordinate(float(sp.X) * float(t))
	v.Y += Coordinate(float(sp.Y) * float(t))
	return v
}
