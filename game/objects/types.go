package objects

type ObjectID int64

type ObjectType string

func (t ObjectType) String() string {
	return string(t)
}
