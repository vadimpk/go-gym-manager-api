package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
)

const (
	errNotInDB    = "sql: no rows in result set"
	errBadRequest = "bad request"
)

type Services struct {
	Managers
	Members
	Memberships
	Trainers
}

type Managers interface {
	SignIn(input domain.SignInInput) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
}

type Members interface {
	CreateNew(input domain.MemberCreateInput) (int, error)
	GetByID(id int) (domain.Member, error)
	GetByPhoneNumber(num string) (domain.Member, error)
	UpdateByID(id int, input domain.MemberUpdateInput) error
	DeleteByID(id int) error
	SetMembership(memberID int, membershipID int) error
	GetMembership(memberID int) (domain.MembersMembershipResponse, error)
	DeleteMembership(memberID int) error
}

type Memberships interface {
	CreateNew(input domain.MembershipCreateInput) (int, error)
	GetByID(id int) (domain.Membership, error)
	UpdateByID(id int, input domain.MembershipUpdateInput) error
	DeleteByID(id int) error
}

type Trainers interface {
	CreateNew(input domain.TrainerCreateInput) (int, error)
	GetByID(id int) (domain.Trainer, error)
	UpdateByID(id int, input domain.TrainerUpdateInput) error
	DeleteByID(id int) error
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewServices(cfg *config.Config, tokenManager auth.TokenManager, repos *repository.Repositories) *Services {
	managerService := NewManagerService(repos.Managers, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	membersService := NewMembersService(repos.Members, repos.Memberships)
	membershipsService := NewMembershipsService(repos.Memberships)
	trainersService := NewTrainersService(repos.Trainers)
	return &Services{
		Managers:    managerService,
		Members:     membersService,
		Memberships: membershipsService,
		Trainers:    trainersService,
	}
}
