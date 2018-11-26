package rest

//easyjson:json
type WebSocketRequest struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

//easyjson:json
type WebSocketResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

//easyjson:json
type AuthRequestData struct {
	AuthToken string `json:"auth_token"`
	UserAgent string `json:"user_agent"`
}

//easyjson:json
type RoomResponseData struct {
}

//easyjson:json
type OkResponseData struct {
	Message string `json:"message"`
}

//easyjson:json
type ErrorResponseData struct {
	Message string `json:"message"`
}
