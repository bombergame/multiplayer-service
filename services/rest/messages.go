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
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

//easyjson:json
type AuthRequestData struct {
	AuthToken string `json:"auth_token"`
	UserAgent string `json:"user_agent"`
}

//easyjson:json
type RoomDataResponse struct {
}

//easyjson:json
type ErrorResponse struct {
	Message string `json:"message"`
}
