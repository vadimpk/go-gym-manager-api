package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository/postgres"
)

type Repositories struct {
	Managers
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Managers: postgres.NewManagerRepo(db),
	}
}

type Managers interface {
	GetByCredentials(ctx context.Context, input domain.SignInInput) (domain.Manager, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Manager, error)
	SetSession(ctx context.Context, managerID int, session domain.Session) error
}
