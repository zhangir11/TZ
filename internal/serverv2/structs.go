package server

import (
	service "authentication/pkg/servicev2"
)

type Handler struct {
	authService service.Service
}

func NewHandler(authService service.Service) *Handler {
	return &Handler{
		authService: authService,
	}
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
