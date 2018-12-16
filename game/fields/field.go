package fields

import (
	"fmt"
	"github.com/bombergame/multiplayer-service/game/cache"
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/bombs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/objects/walls/solid"
	"github.com/bombergame/multiplayer-service/game/objects/walls/weak"
	"github.com/bombergame/multiplayer-service/game/physics"
	"log"
	"math/rand"
	"time"
)

type Field struct {
	size physics.Size2D

	bCache *cache.Queue

	objects    [][]objects.GameObject
	explosives [][]objects.ExplosiveObject

	invalidCellIndexError *errs.InvalidCellIndexError
}

func NewField(size physics.Size2D) *Field {
	f := &Field{
		size: size,

		explosives: func() [][]objects.ExplosiveObject {
			b := make([][]objects.ExplosiveObject, size.Height)
			for i := physics.Integer(0); i < size.Height; i++ {
				b[i] = make([]objects.ExplosiveObject, size.Width)
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
			log.Println("Checking position: ", pos)

			x, y := posToInt(pos)
			if x < 0 || x >= f.size.Width || y < 0 || y >= f.size.Height {
				log.Println("Cannot move. Cell index out of range")
				return nil, f.invalidCellIndexError
			}
			if f.explosives[y][x] != nil {
				log.Println("Cannot move. Explosive")
				if b, ok := f.explosives[y][x].(*bombs.Bomb); ok {
					log.Println("Cannot move. Explosive is a bomb")
					return b, nil
				}
			}

			log.Println("Check if can move: ", f.objects[y][x])
			return f.objects[y][x], nil
		})
		p.SetMovementHandler(func(pOld, pNew physics.PositionVec2D) {
			log.Println("Before move:")
			f.print()

			xOld, yOld := posToInt(pOld)
			xNew, yNew := posToInt(pNew)
			obj := f.objects[yOld][xOld]
			f.objects[yOld][xOld] = nil
			f.objects[yNew][xNew] = obj

			log.Println("After move:")
			f.print()
		})
		p.SetDropBombHandler(func(pos physics.PositionVec2D) {
			x, y := posToInt(pos)
			if f.explosives[y][x] != nil {
				return
			}

			v, _ := f.bCache.Dequeue()
			b := v.(*bombs.Bomb)
			b.Spawn(pos)
			f.explosives[y][x] = b

			log.Println("Bomb placed:", pos)
		})

		pArr = append(pArr, p)
	}

	index := 0
	d := f.size.Height / n

	rand.Seed(time.Now().UTC().UnixNano())
	for y := physics.Integer(0); y < f.size.Height; y++ {
		if y%d == 0 {
			x := rand.Intn(int(f.size.Width))
			f.objects[y][x] = pArr[index]
			index++
		} else {
			for x := physics.Integer(1); x < f.size.Width-1; x++ {
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
	}
}

const (
	EmptyProb     = 0.5
	WeakWallProb  = 0.7
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
		}
	}

	f.bCache = cache.NewQueue()
	for i := physics.Integer(0); i < f.size.Width*f.size.Height; i++ {
		b := bombs.NewBomb()

		objID++
		b.SetObjectID(objID)
		b.SetObjectType(bombs.Type)

		b.SetChangeHandler(h)
		b.SetExplosionHandler(func(obj objects.ExplosiveObject) {
			log.Println("Explosion: ", obj)

			_ = f.bCache.Enqueue(obj)

			x, y := posToInt(obj.Transform().Position)
			f.explosives[y][x] = nil

			d := obj.ExplosionRadius()
			for i := physics.Integer(0); i < d; i++ {
				f.destroyObject(x-i, y)
				f.destroyObject(x+i, y)
				f.destroyObject(x, y-i)
				f.destroyObject(x, y+i)
			}
		})

		f.bCache.Add(b)
	}
}

func (f *Field) UpdateObjects(d time.Duration) {
	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			if f.explosives[i][j] != nil {
				f.explosives[i][j].Update(d)
			}
			if f.objects[i][j] != nil {
				f.objects[i][j].Update(d)
			}
		}
	}
}

func (f *Field) destroyObject(x, y physics.Integer) {
	if x < 0 || x >= f.size.Width || y < 0 || y >= f.size.Height {
		return
	}
	if f.objects[y][x] != nil {
		if obj, ok := f.objects[y][x].(objects.DestructableObject); ok {
			obj.Collapse()
			f.objects[y][x] = nil
		}
	} else if f.explosives[y][x] != nil {
		if obj, ok := f.explosives[y][x].(objects.DestructableObject); ok {
			obj.Collapse()
			f.explosives[y][x] = nil
		}
	}
}

func posToInt(p physics.PositionVec2D) (physics.Integer, physics.Integer) {
	return physics.Integer(p.X), physics.Integer(p.Y)
}

func (f *Field) print() {
	for i := physics.Integer(0); i < f.size.Height; i++ {
		for j := physics.Integer(0); j < f.size.Width; j++ {
			if f.explosives[i][j] != nil {
				fmt.Print("o")
			} else if f.objects[i][j] != nil {
				if f.objects[i][j].ObjectType() == players.Type {
					fmt.Print("P")
				} else if f.objects[i][j].ObjectType() == weakwalls.Type {
					fmt.Print("+")
				} else if f.objects[i][j].ObjectType() == solidwalls.Type {
					fmt.Print("#")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}