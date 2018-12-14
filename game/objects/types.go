package objects

type ID int64

type ObjectType string

func (t ObjectType) String() string {
	return string(t)
}
