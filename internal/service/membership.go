package service

import (
	"errors"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"time"
)

type MembershipsService struct {
	repo repository.Memberships
}

func NewMembershipsService(repo repository.Memberships) *MembershipsService {
	return &MembershipsService{repo: repo}
}

func (s *MembershipsService) CreateNew(input domain.MembershipCreateInput) (int, error) {
	_, err := time.ParseDuration(input.Duration)
	if err != nil {
		return 0, errors.New(errBadRequest)
	}
	return s.repo.Create(input)
}

func (s *MembershipsService) GetByID(id int) (domain.Membership, error) {
	return s.repo.GetByID(id)
}

func (s *MembershipsService) UpdateByID(id int, input domain.MembershipUpdateInput) error {
	// check if exists
	membership, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// this is done in case not all information is provided in json input
	// not sure if this is needed, maybe this should be done on frontend side by using default placeholders
	if input.ShortName == "" {
		input.ShortName = membership.ShortName
	}
	if input.Description == "" {
		input.Description = membership.Description
	}
	if input.Price == 0 {
		input.Price = membership.Price
	}
	if input.Duration == "" {
		input.Duration = membership.Duration
	} else {
		if _, err := time.ParseDuration(input.Duration); err != nil {
			return errors.New(errBadRequest)
		}
	}
	return s.repo.Update(id, input)
}

func (s *MembershipsService) DeleteByID(id int) error {
	// check if exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
