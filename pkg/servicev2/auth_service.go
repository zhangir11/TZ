package service

import (
	"authentication/config"
	storage "authentication/pkg/storagev2"
	"authentication/pkg/utils"
	"time"
)

//-------------------------------------------------------------------------------------------------------

func (a *AuthService) CreateAuthSession(guid string) (*AuthSession, int) {
	if ok := utils.IsValidGuid(guid); !ok {
		return &AuthSession{}, 400
	}

	if _, err := a.authStorage.Get(guid); err == nil {
		return &AuthSession{}, 401
	}

	accessToken, err := utils.NewAccessToken(guid)
	if err != nil {
		return &AuthSession{}, 500
	}

	refreshToken, err := utils.NewRefreshToken(accessToken.Token)
	if err != nil {
		return &AuthSession{}, 500
	}

	refreshTokenHash, err := utils.HashToken(refreshToken)
	if err != nil {
		return &AuthSession{}, 500
	}

	timeDelta, err := time.ParseDuration(config.Conf.RefreshTokenTimeDelta)
	if err != nil {
		return &AuthSession{}, 500
	}

	if err := a.authStorage.Insert(&storage.Session{
		Guid:         guid,
		RefreshToken: refreshTokenHash,
		ExpiredAt:    time.Now().Add(timeDelta).UTC().Unix(),
	}); err != nil {
		return &AuthSession{}, 500
	}

	return &AuthSession{
		AccessToken:  accessToken.Token,
		ExpiresAt:    accessToken.Claims.ExpiresAt,
		RefreshToken: refreshToken,
	}, 201
}

//-------------------------------------------------------------------------------------------------------

func (a *AuthService) DeleteAuthSession(accessToken, refreshToken string) (string, int) {
	if ok := utils.CompareRefreshAndAccessToken(accessToken, refreshToken); !ok {
		return "", 400
	}

	accessTokenData, err := utils.ParseAccessToken(accessToken)
	if err != nil {
		return "", 400
	}

	if ok := utils.IsExpired(accessTokenData.Claims.ExpiresAt); !ok {
		return "", 401
	}

	session, err := a.authStorage.Get(accessTokenData.Claims.Id)
	if err != nil {
		return "", 401
	}

	if ok := utils.CompareHashAndToken(refreshToken, session.RefreshToken); !ok {
		return "", 401
	}

	if ok := utils.IsExpired(session.ExpiredAt); ok {
		return "", 401
	}

	if err := a.authStorage.Delete(accessTokenData.Claims.Id); err != nil {
		return "", 500
	}

	return accessTokenData.Claims.Id, 200
}
