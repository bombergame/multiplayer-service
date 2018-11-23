package field

import (
	"github.com/bombergame/multiplayer-service/game/objects/field/cell"
	"github.com/bombergame/multiplayer-service/game/physics"
	"math"
)

type Field struct {
	size  Size
	cells [][]*cell.Cell
}

func NewField(s Size) *Field {
	return &Field{
		size:  s,
		cells: createCells(s),
	}
}

func createCells(s Size) [][]*cell.Cell {
	cells := make([][]*cell.Cell, s.Height)
	for i := 0; i < int(s.Height); i++ {
		cells[i] = make([]*cell.Cell, s.Width)
	}

	r := cell.NewObjectTypeRandomizerImpl()

	for i := 0; i < int(s.Height); i++ {
		for j := 0; j < int(s.Width); j++ {
			cells[i][j] = cell.NewCell()
			cells[i][j].SpawnRandomObject(r)
		}
	}

	return cells
}

func (f *Field) IsValidPosition(p physics.PositionVec2D) bool {
	xInt, yInt := positionVec2DToCellIndexes(p)
	return xInt >= 0 && xInt < f.size.Width &&
		yInt >= 0 && yInt < f.size.Height
}

func (f *Field) IsCellEmpty(p physics.PositionVec2D) bool {
	xInt, yInt := positionVec2DToCellIndexes(p)
	return f.cells[xInt][yInt].IsEmpty()
}

func positionVec2DToCellIndexes(p physics.PositionVec2D) (Integer, Integer) {
	return Integer(math.Round(float64(p.X))), Integer(math.Round(float64(p.Y)))
}
