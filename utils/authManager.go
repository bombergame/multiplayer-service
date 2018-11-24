package utils

import (
	"github.com/bombergame/common/auth"
	"math/rand"
)

type AuthManager struct {
}

func NewAuthManager() *AuthManager {
	return &AuthManager{}
}

func (m *AuthManager) GetProfileInfo(authToken string, userAgent string) (*auth.ProfileInfo, error) {
	//TODO
	return &auth.ProfileInfo{ID: rand.Int63()}, nil
}
