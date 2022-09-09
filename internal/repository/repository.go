package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository/postgres"
)

type Repositories struct {
	Managers
	Members
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Managers: postgres.NewManagerRepo(db),
		Members:  postgres.NewMemberRepo(db),
	}
}

type Managers interface {
	GetByCredentials(input domain.SignInInput) (domain.Manager, error)
	GetByRefreshToken(refreshToken string) (int, error)
	SetSession(managerID int, session domain.Session) error
}

type Members interface {
	Create(input domain.MemberCreate) (int, error)
	GetByID(id int) (domain.Member, error)
	GetByPhoneNumber(num string) (domain.Member, error)
	Update(id int, input domain.MemberUpdate) error
	Delete(id int) error
	SetMembership(id int, membershipID int) error
	DeleteMembership(id int) error
}
