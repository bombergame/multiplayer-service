package main

import (
	"encoding/json"
	"fmt"
	"github.com/bombergame/multiplayer-service/game/objects/wall"
)

func main() {
	w := wall.NewWall()
	b, _ := json.Marshal(w)
	fmt.Println(string(b))
}
