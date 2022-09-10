package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository/postgres"
	"time"
)

type Repositories struct {
	Managers
	Members
	Memberships
	Trainers
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Managers:    postgres.NewManagerRepo(db),
		Members:     postgres.NewMemberRepo(db),
		Memberships: postgres.NewMembershipRepo(db),
		Trainers:    postgres.NewTrainerRepo(db),
	}
}

type Managers interface {
	GetByCredentials(input domain.SignInInput) (domain.Manager, error)
	GetByRefreshToken(refreshToken string) (int, error)
	SetSession(managerID int, session domain.Session) error
}

type Members interface {
	Create(input domain.MemberCreateInput) (int, error)
	GetByID(id int) (domain.Member, error)
	GetByPhoneNumber(num string) (domain.Member, error)
	Update(id int, input domain.MemberUpdateInput) error
	Delete(id int) error
	SetMembership(memberID int, membershipID int, expiresAt time.Time) error
	UpdateMembership(memberID int, membershipID int, expiresAt time.Time) error
	GetMembership(memberID int) (domain.MembersMembership, error)
	DeleteMembership(memberID int) error
}

type Memberships interface {
	Create(input domain.MembershipCreateInput) (int, error)
	GetByID(id int) (domain.Membership, error)
	Update(id int, input domain.MembershipUpdateInput) error
	Delete(id int) error
}

type Trainers interface {
	Create(input domain.TrainerCreateInput) (int, error)
	GetByID(id int) (domain.Trainer, error)
	Update(id int, input domain.TrainerUpdateInput) error
	Delete(id int) error
}
