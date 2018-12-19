package auth

//AuthenticationManager provides user authentication method
type AuthenticationManager interface {
	GetProfileInfo(authToken string, userAgent string) (*ProfileInfo, error)
}
