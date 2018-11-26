package rest

//easyjson:json
type WebSocketMessage struct {
	Type string `json:"type"`
	Data []byte `data:"data"`
}

//easyjson:json
type AuthRequestData struct {
	AuthToken string `json:"auth_token"`
}

//easyjson:json
type ErrorResponse struct {
	Message string `json:"message"`
}
