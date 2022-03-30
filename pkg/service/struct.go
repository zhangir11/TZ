package service

import "authentication/pkg/storage"

type AuthSession struct {
	AccessToken  string
	ExpiresAt    int64
	RefreshToken string
}

type AuthService struct {
	authStorage storage.Manager
}

func NewAuthService(authStorage storage.Manager) *AuthService {
	return &AuthService{
		authStorage: authStorage,
	}
}
