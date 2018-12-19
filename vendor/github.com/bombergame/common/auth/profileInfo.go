package auth

//go:generate easyjson

//ProfileInfo contains authorized user info
//easyjson:json
type ProfileInfo struct {
	ID int64 `json:"profile_id"`
}
