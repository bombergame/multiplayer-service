package cell

import (
	"math/rand"
)

type ObjectType string

const (
	Empty     = "field.cell.empty"
	WallWeak  = "field.cell.wall.weak"
	WallSolid = "field.cell.wall.solid"
)

type ObjectTypeRandomizer interface {
	GetObjectType() ObjectType
}

type ObjectTypeRandomizerImpl struct {
	pEmpty     float32
	pWallSolid float32
	pWallWeak  float32
}

func NewObjectTypeRandomizerImpl() *ObjectTypeRandomizerImpl {
	return &ObjectTypeRandomizerImpl{
		pEmpty:     0.00,
		pWallSolid: 0.60,
		pWallWeak:  0.80,
	}
}

func (r *ObjectTypeRandomizerImpl) GetObjectType() ObjectType {
	p := rand.Float32()

	if p > r.pWallWeak {
		return WallWeak
	}

	if p > r.pWallSolid {
		return WallSolid
	}

	return Empty
}
