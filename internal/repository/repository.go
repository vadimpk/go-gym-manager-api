package repository

import (
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
	GetByCredentials(input domain.SignInInput) (domain.Manager, error)
	GetByRefreshToken(refreshToken string) (int, error)
	SetSession(managerID int, session domain.Session) error
}
