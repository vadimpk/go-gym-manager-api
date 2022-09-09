package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
)

type Services struct {
	Managers
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewServices(cfg *config.Config, tokenManager auth.TokenManager, repos *repository.Repositories) *Services {
	managerService := NewManagerService(repos.Managers, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	return &Services{Managers: managerService}
}

type Managers interface {
	SignIn(input domain.SignInInput) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
}
