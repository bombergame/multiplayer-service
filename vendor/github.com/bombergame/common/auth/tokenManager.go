package auth

import (
	"time"
)

const (
	//ExpireTimeFormat is an alias for time.RFC3339
	ExpireTimeFormat = time.RFC3339

	//DefaultTokenValidDuration is the default value of time the token is valid
	DefaultTokenValidDuration = 15 * time.Minute
)

//TokenInfo contains data encrypted in token
type TokenInfo struct {
	ProfileID  string `mapstructure:"profile_id"`
	UserAgent  string `mapstructure:"user_agent"`
	ExpireTime string `mapstructure:"expire_time"`
}

//TokenManager provides methods to manage access tokens
type TokenManager interface {
	CreateToken(info TokenInfo) (string, error)
	ParseToken(token string) (*TokenInfo, error)
}
