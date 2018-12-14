package fields

import (
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/objects/walls/solid"
	"github.com/bombergame/multiplayer-service/game/objects/walls/weak"
	"github.com/bombergame/multiplayer-service/game/physics"
	"math/rand"
	"time"
)

type Field struct {
	size physics.Size2D

	cells                 [][]objects.GameObject
	invalidCellIndexError *errs.InvalidCellIndexError
}

func NewField(size physics.Size2D) *Field {
	f := &Field{
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

	return f
}

func (f *Field) PlacePlayers(pAll map[int64]*players.Player) {
	x, y := physics.Integer(0), physics.Integer(0)
	for _, p := range pAll {
		if x == f.size.Width {
			y++
		}
		if y == f.size.Height {
			break
		}

		p.SetObjectType(players.Type)
		p.SetCellObjectGetter(func(pos physics.PositionVec2D) (objects.GameObject, *errs.InvalidCellIndexError) {
			x, y := physics.Integer(pos.X), physics.Integer(pos.Y)
			if x < 0 || x >= f.size.Width || y < 0 || y >= f.size.Height {
				return nil, f.invalidCellIndexError
			}
			return f.cells[y][x], nil
		})

		f.cells[y][x] = p
		x++
	}
}

const (
	EmptyProb     = 0.5
	WeakWallProb  = 0.6
	SolidWallProb = 1.0
)

func (f *Field) SpawnObjects(h objects.ChangeHandler) {
	objID := objects.ObjectID(0)

	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			var obj objects.GameObject

			if f.cells[i][j] == nil {
				prob := rand.NormFloat64()

				if prob < EmptyProb {
					f.cells[i][j] = nil
					continue
				}

				if prob < WeakWallProb {
					obj = weakwalls.NewWall()
					obj.SetObjectType(weakwalls.Type)
				} else {
					obj = solidwalls.NewWall()
					obj.SetObjectType(solidwalls.Type)
				}
			} else {
				obj = f.cells[i][j]
			}

			objID++
			obj.SetObjectID(objID)
			obj.SetChangeHandler(h)
			obj.Spawn(physics.GetPositionVec2D(physics.Coordinate(i), physics.Coordinate(j)))

			f.cells[i][j] = obj
		}
	}
}

func (f *Field) UpdateObjects(d time.Duration) {
	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			if f.cells[i][j] == nil {
				continue
			}
			f.cells[i][j].Update(d)
		}
	}
}
