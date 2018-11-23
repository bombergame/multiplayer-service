package field

type Integer int8

type Size struct {
	Width  Integer
	Height Integer
}

const (
	DefaultWidth  = 10
	DefaultHeight = 10
)

func GetSize(w, h Integer) Size {
	return Size{Width: w, Height: h}
}
