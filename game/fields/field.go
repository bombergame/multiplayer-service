package fields

import (
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/walls"
	"math"
	"math/rand"
)

type Size struct {
	Width  int32
	Height int32
}

type Field struct {
	size  Size
	cells [][]objects.GameObject
}

func NewField(size Size) *Field {
	return &Field{
		size: size,
		cells: func() [][]objects.GameObject {
			c := make([][]objects.GameObject, size.Height)
			var i int32
			for i = 0; i < size.Height; i++ {
				c[i] = make([]objects.GameObject, size.Width)
			}
			return c
		}(),
	}
}

const (
	SolidWallsPercent = 0.5
	WeakWallsPercent  = 0.2
)

func (f *Field) GenerateRandom(nPlayers int32) {
	var i, j int32
	for i = 0; i < f.size.Height; i++ {
		for j = 0; j < f.size.Width; j++ {
			f.cells[i][j] = nil
		}
	}

	n := (f.size.Width * f.size.Height) - nPlayers

	nSolidWalls := f.countNumObjects(n, SolidWallsPercent)
	for i = 0; i < nSolidWalls; i++ {
		r, c := f.randIndexes()
		f.cells[r][c] = &walls.SolidWall{}
	}

	nWeakWalls := f.countNumObjects(n, WeakWallsPercent)
	for i = 0; i < nWeakWalls; i++ {
		r, c := f.randIndexes()
		f.cells[r][c] = &walls.SolidWall{}
	}
}

func (f *Field) randIndexes() (int, int) {
	return f.randIndex(f.size.Height), f.randIndex(f.size.Width)
}

func (f *Field) randIndex(maxIndex int32) int {
	return int(rand.Int31n(maxIndex))
}

func (f *Field) countNumObjects(n int32, p float64) int32 {
	return int32(math.Floor(float64(n) * p))
}
