package cell

type Cell struct {
	objectType ObjectType
}

func NewCell() *Cell {
	return &Cell{
		objectType: Empty,
	}
}

func (c *Cell) GetObjectType() ObjectType {
	return c.objectType
}

func (c *Cell) IsEmpty() bool {
	return c.objectType == Empty
}

func (c *Cell) SpawnRandomObject(r ObjectTypeRandomizer) {
	c.objectType = r.GetObjectType()
}
