package rest

import (
	"github.com/mailru/easyjson"
)

type Serializable interface {
	easyjson.Marshaler
	easyjson.Unmarshaler
}

//easyjson:json
type WebSocketMessage struct {
	Type string `json:"type"`
	Data string `data:"data"`
}

//easyjson:json
type AuthRequestData struct {
	AuthToken string `json:"auth_token"`
}

//easyjson:json
type ErrorResponse struct {
	Message string `json:"message"`
}
