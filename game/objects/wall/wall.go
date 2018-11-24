package wall

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/mailru/easyjson/jwriter"
)

type Wall struct {
	objType   objects.ObjectType
	transform transform.Transform
	collider  collider.Collider
}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) Collider() collider.Collider {
	return w.collider
}

func (w *Wall) Start() {
	//TODO
}

func (w *Wall) Update(timeDiff physics.Time) {
	//TODO
}

func (w *Wall) spawn() {
	//TODO
}

//easyjson:json
type WallJSON struct {
	ObjType   objects.ObjectType  `json:"object_type"`
	Transform transform.Transform `json:"transform"`
	Collider  collider.Collider   `json:"collider"`
}

func (w *Wall) MarshalEasyJSON(wr *jwriter.Writer) {
	wJSON := &WallJSON{
		ObjType:   w.objType,
		Transform: w.transform,
		Collider:  w.collider,
	}
	wJSON.MarshalEasyJSON(wr)
}
