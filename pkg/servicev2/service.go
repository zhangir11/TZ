package service

type Service interface {
	CreateAuthSession(guid string) (*AuthSession, int)
	DeleteAuthSession(accessToken, refreshToken string) (string, int)
}
