package physics

//go:generate easyjson

//easyjson:json
type Size2D struct {
	Width  Integer `json:"width"`
	Height Integer `json:"height"`
}

func GetSize2D(w, h Integer) Size2D {
	return Size2D{
		Width:  w,
		Height: h,
	}
}
