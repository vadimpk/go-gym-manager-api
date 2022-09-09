package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
	"strconv"
	"time"
)

type ManagerService struct {
	repo         repository.Managers
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewManagerService(repo repository.Managers, tokenManager auth.TokenManager, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *ManagerService {
	return &ManagerService{repo: repo, tokenManager: tokenManager, accessTokenTTL: accessTokenTTL, refreshTokenTTL: refreshTokenTTL}
}

func (s *ManagerService) SignIn(input domain.SignInInput) (Tokens, error) {
	manager, err := s.repo.GetByCredentials(input)
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(manager.ID)
}

func (s *ManagerService) RefreshTokens(refreshToken string) (Tokens, error) {
	managerID, err := s.repo.GetByRefreshToken(refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(managerID)
}

func (s *ManagerService) createSession(managerID int) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(strconv.Itoa(managerID), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.repo.SetSession(managerID, session)

	return res, err
}
