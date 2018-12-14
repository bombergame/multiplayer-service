package ws

//go:generate easyjson

//easyjson:json
type InMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

//easyjson:json
type OutMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type OutChan chan OutMessage

const (
	OutChanLen = 10

	AuthMessageType = "auth"

	RoomMessageType   = "room"
	FieldMessageType  = "field"
	TickerMessageType = "ticker"

	OkMessageType    = "ok"
	ErrorMessageType = "error"
)

//easyjson:json
type AuthMessageData struct {
	AuthToken string `json:"auth_token"`
	UserAgent string `json:"user_agent"`
}

//easyjson:json
type OkMessageData struct {
	Message string `json:"message"`
}

//easyjson:json
type ErrorMessageData struct {
	Message string `json:"message"`
}
