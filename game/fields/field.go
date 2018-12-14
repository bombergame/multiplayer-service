package fields

import (
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/objects/walls/solid"
	"github.com/bombergame/multiplayer-service/game/objects/walls/weak"
	"github.com/bombergame/multiplayer-service/game/physics"
	"math/rand"
)

type Field struct {
	size  physics.Size2D
	cells [][]objects.GameObject

	invalidCellIndexError *errs.InvalidCellIndexError
}

func NewField(size physics.Size2D) *Field {
	return &Field{
		size: size,
		cells: func() [][]objects.GameObject {
			c := make([][]objects.GameObject, size.Height)
			for i := physics.Integer(0); i < size.Height; i++ {
				c[i] = make([]objects.GameObject, size.Width)
			}
			return c
		}(),

		invalidCellIndexError: errs.NewInvalidCellIndexError(),
	}
}

func (f *Field) SpawnPlayers(pAll map[int64]*players.Player) {
	x, y := physics.Integer(0), physics.Integer(0)
	for _, p := range pAll {
		if x == f.size.Width {
			y++
		}
		if y == f.size.Height {
			break
		}

		f.cells[x][y] = p
		x++
	}
}

const (
	EmptyProb     = 0.5
	WeakWallProb  = 0.6
	SolidWallProb = 1.0
)

func (f *Field) SpawnObjects(h objects.ChangeHandler) {
	objID := objects.ID(0)

	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			objID++

			if f.cells[i][j] != nil {
				continue
			}

			prob := rand.NormFloat64()

			if prob < EmptyProb {
				f.cells[i][j] = nil
				continue
			}

			var obj objects.GameObject
			if prob < WeakWallProb {
				obj = weakwalls.NewWall()
			} else {
				obj = solidwalls.NewWall()
			}

			obj.SetObjectID(objID)
			obj.SetChangeHandler(h)
			obj.Spawn(physics.GetPositionVec2D(physics.Coordinate(i), physics.Coordinate(j)))

			f.cells[i][j] = obj
		}
	}
}
