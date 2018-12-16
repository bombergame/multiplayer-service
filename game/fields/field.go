package fields

import (
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/bombs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/objects/walls/solid"
	"github.com/bombergame/multiplayer-service/game/objects/walls/weak"
	"github.com/bombergame/multiplayer-service/game/physics"
	"math/rand"
	"time"
)

type Field struct {
	size physics.Size2D

	bombs   [][]*bombs.Bomb
	objects [][]objects.GameObject

	invalidCellIndexError *errs.InvalidCellIndexError
}

func NewField(size physics.Size2D) *Field {
	f := &Field{
		size: size,

		bombs: func() [][]*bombs.Bomb {
			b := make([][]*bombs.Bomb, size.Height)
			for i := physics.Integer(0); i < size.Height; i++ {
				b[i] = make([]*bombs.Bomb, size.Width)
			}
			return b
		}(),
		objects: func() [][]objects.GameObject {
			o := make([][]objects.GameObject, size.Height)
			for i := physics.Integer(0); i < size.Height; i++ {
				o[i] = make([]objects.GameObject, size.Width)
			}
			return o
		}(),

		invalidCellIndexError: errs.NewInvalidCellIndexError(),
	}

	return f
}

func (f *Field) Size() physics.Size2D {
	return f.size
}

func (f *Field) PlaceObjects(pAll map[int64]*players.Player) {
	n := physics.Integer(len(pAll))
	pArr := make([]*players.Player, 0, n)

	for _, p := range pAll {
		p.SetObjectType(players.Type)
		p.SetCellObjectGetter(func(pos physics.PositionVec2D) (objects.GameObject, *errs.InvalidCellIndexError) {
			x, y := physics.Integer(pos.X), physics.Integer(pos.Y)
			if x < 0 || x >= f.size.Width || y < 0 || y >= f.size.Height {
				return nil, f.invalidCellIndexError
			}
			return f.objects[y][x], nil
		})
		p.SetMovementHandler(func(pOld, pNew physics.PositionVec2D) {
			xOld, yOld := physics.Integer(pOld.X), physics.Integer(pOld.Y)
			xNew, yNew := physics.Integer(pNew.X), physics.Integer(pNew.Y)
			obj := f.objects[yOld][xOld]
			f.objects[yOld][xOld] = nil
			f.objects[yNew][xNew] = obj
		})

		pArr = append(pArr, p)
	}

	index := 0
	d := f.size.Height / n

	for y := physics.Integer(0); y < f.size.Height; y++ {
		if y%d == 0 {
			x := rand.Intn(int(f.size.Width))
			f.objects[y][x] = pArr[index]
			index++
		} else {
			for x := physics.Integer(1); x < f.size.Width-2; x++ {
				prob := rand.NormFloat64()

				if prob < EmptyProb {
					f.objects[y][x] = nil
					continue
				}

				var obj objects.GameObject
				if prob < WeakWallProb {
					obj = weakwalls.NewWall()
					obj.SetObjectType(weakwalls.Type)
				} else {
					obj = solidwalls.NewWall()
					obj.SetObjectType(solidwalls.Type)
				}
				f.objects[y][x] = obj
			}
		}

		for x := physics.Integer(0); x < f.size.Width; x++ {
			var b *bombs.Bomb
			b = bombs.NewBomb()
			b.SetObjectType(bombs.Type)
			f.bombs[y][x] = b
		}
	}
}

const (
	EmptyProb     = 0.5
	WeakWallProb  = 0.6
	SolidWallProb = 1.0
)

func (f *Field) SpawnObjects(h objects.ChangeHandler) {
	objID := objects.ObjectID(0)

	for y := physics.Integer(0); y < f.size.Height; y++ {
		for x := physics.Integer(0); x < f.size.Width; x++ {
			if f.objects[y][x] == nil {
				continue
			}

			obj := f.objects[y][x]
			objID++
			obj.SetObjectID(objID)
			obj.SetChangeHandler(h)
			obj.Spawn(physics.GetPositionVec2D(physics.Coordinate(y), physics.Coordinate(x)))

			b := f.bombs[y][x]
			objID++
			b.SetObjectID(objID)
			b.SetChangeHandler(h)
		}
	}
}

func (f *Field) UpdateObjects(d time.Duration) {
	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			if f.objects[i][j] == nil {
				continue
			}
			f.objects[i][j].Update(d)
		}
	}
}
