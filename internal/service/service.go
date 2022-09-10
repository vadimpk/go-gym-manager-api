package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
)

type Services struct {
	Managers
	Members
	Memberships
}

type Managers interface {
	SignIn(input domain.SignInInput) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
}

type Members interface {
	CreateNew(input domain.MemberCreate) (int, error)
	GetByID(id int) (domain.Member, error)
	GetByPhoneNumber(num string) (domain.Member, error)
	UpdateByID(id int, input domain.MemberUpdate) error
	DeleteByID(id int) error
	SetMembership(id int, membershipID int) error
	DeleteMembership(id int) error
}

type Memberships interface {
	CreateNew(input domain.MembershipCreateInput) (int, error)
	GetByID(id int) (domain.Membership, error)
	UpdateByID(id int, input domain.MembershipUpdateInput) error
	DeleteByID(id int) error
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewServices(cfg *config.Config, tokenManager auth.TokenManager, repos *repository.Repositories) *Services {
	managerService := NewManagerService(repos.Managers, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	membersService := NewMembersService(repos.Members)
	membershipsService := NewMembershipsService(repos.Memberships)
	return &Services{
		Managers:    managerService,
		Members:     membersService,
		Memberships: membershipsService,
	}
}
