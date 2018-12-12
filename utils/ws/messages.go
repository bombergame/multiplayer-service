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

//easyjson:json
type AuthMessageData struct {
	AuthToken string `json:"auth_token"`
	UserAgent string `json:"user_agent"`
}

//easyjson:json
type RoomMessageData struct {
	Title   string  `json:"title"`
	State   string  `json:"state"`
	Players []int64 `json:"players"`
}

//easyjson:json
type OkMessageData struct {
	Message string `json:"message"`
}

//easyjson:json
type ErrorMessageData struct {
	Message string `json:"message"`
}

type InChan chan InMessage
type OutChan chan OutMessage
