package service

import (
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"time"
)

type MembersService struct {
	repo            repository.Members
	membershipsRepo repository.Memberships
}

func NewMembersService(repo repository.Members, membershipsRepo repository.Memberships) *MembersService {
	return &MembersService{repo: repo, membershipsRepo: membershipsRepo}
}

func (s *MembersService) CreateNew(input domain.MemberCreateInput) (int, error) {
	return s.repo.Create(input)
}

func (s *MembersService) GetByID(id int) (domain.Member, error) {
	return s.repo.GetByID(id)
}

func (s *MembersService) GetByPhoneNumber(num string) (domain.Member, error) {
	return s.repo.GetByPhoneNumber(num)
}

func (s *MembersService) UpdateByID(id int, input domain.MemberUpdateInput) error {
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

func (s *MembersService) SetMembership(memberID int, membershipID int) error {
	// check if membership exists
	membershipNew, err := s.membershipsRepo.GetByID(membershipID)
	if err != nil {
		return err
	}

	duration, _ := time.ParseDuration(membershipNew.Duration)
	expiresAt := time.Now().Add(duration)

	// check if member has active membership
	_, err = s.repo.GetMembership(memberID)
	if err != nil {
		if err.Error() == errNotInDB {
			return s.repo.SetMembership(memberID, membershipID, expiresAt)
		}
		return err
	}
	// TODO maybe setting new membership shouldn't be allowed if current membership hasn't expired yet
	return s.repo.UpdateMembership(memberID, membershipID, expiresAt)
}

func (s *MembersService) DeleteMembership(memberID int) error {
	// check if member has active membership
	_, err := s.repo.GetMembership(memberID)
	if err != nil {
		return err
	}
	return s.repo.DeleteMembership(memberID)
}
