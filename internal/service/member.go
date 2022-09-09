package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
)

type MembersService struct {
	repo repository.Members
}

func NewMembersService(repo repository.Members) *MembersService {
	return &MembersService{repo: repo}
}

func (s *MembersService) CreateNew(input domain.MemberCreate) (int, error) {
	// TODO check if membership id exists
	return s.repo.Create(input)
}

func (s *MembersService) GetByID(id int) (domain.Member, error) {
	return s.repo.GetByID(id)
}

func (s *MembersService) GetByPhoneNumber(num string) (domain.Member, error) {
	return s.repo.GetByPhoneNumber(num)
}

func (s *MembersService) UpdateByID(id int, input domain.MemberUpdate) error {
	member, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// this is done in case not all information is provided in json input
	// not sure if this is needed, maybe this should be done on frontend side by using default placeholders
	if input.FirstName == "" {
		input.FirstName = member.FirstName
	}
	if input.LastName == "" {
		input.LastName = member.LastName
	}
	if input.PhoneNumber == "" {
		input.PhoneNumber = member.PhoneNumber
	}

	return s.repo.Update(id, input)
}

func (s *MembersService) DeleteByID(id int) error {
	// TODO there is probably easier way to do this on repository layer
	// check if exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MembersService) SetMembership(id int, membershipID int) error {
	return s.repo.SetMembership(id, membershipID)
}
func (s *MembersService) DeleteMembership(id int) error {
	return s.repo.DeleteMembership(id)
}
