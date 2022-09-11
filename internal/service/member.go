package service

import (
	"errors"
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
		if err.Error() == domain.ErrNotInDB {
			return s.repo.SetMembership(memberID, membershipID, expiresAt)
		}
		return err
	}
	// TODO maybe setting new membership shouldn't be allowed if current membership hasn't expired yet
	return s.repo.UpdateMembership(memberID, membershipID, expiresAt)
}

func (s *MembersService) GetMembership(memberID int) (domain.MembersMembershipResponse, error) {
	currentMembership, err := s.repo.GetMembership(memberID)
	if err != nil {
		if err.Error() == domain.ErrNotInDB {
			return domain.MembersMembershipResponse{}, errors.New(domain.ErrDoesntHaveMembership)
		}
		return domain.MembersMembershipResponse{}, err
	}
	membershipInfo, err := s.membershipsRepo.GetByID(currentMembership.MembershipID)
	if err != nil {
		return domain.MembersMembershipResponse{}, err
	}

	return domain.MembersMembershipResponse{
		Membership: membershipInfo,
		ExpiresAt:  currentMembership.ExpiresAt,
	}, nil
}

func (s *MembersService) DeleteMembership(memberID int) error {
	// check if member has active membership
	_, err := s.repo.GetMembership(memberID)
	if err != nil {
		return err
	}
	return s.repo.DeleteMembership(memberID)
}

func (s *MembersService) SetNewVisit(memberID int, managerID int) error {
	// check if member exists
	_, err := s.repo.GetByID(memberID)
	if err != nil {
		return err
	}

	// check if membership isn't expired yet
	membership, err := s.GetMembership(memberID)
	if err != nil {
		return err
	}
	if !time.Now().After(membership.ExpiresAt) {
		return errors.New(domain.ErrExpiredMembership)
	}

	// get the latest visit to see if it is possible to set new visit
	visit, err := s.repo.GetLatestVisit(memberID)
	if err != nil {
		if err.Error() == domain.ErrNotInDB {
			return s.repo.SetNewVisit(memberID, managerID)
		}
		return err
	}
	if visit.LeftAt.IsZero() {
		return errors.New(domain.ErrStillInGym)
	}

	return s.repo.SetNewVisit(memberID, managerID)
}

func (s *MembersService) EndVisit(memberID int) error {
	// check if member exists
	_, err := s.repo.GetByID(memberID)
	if err != nil {
		return err
	}

	// get the latest visit to see if it is possible to end the visit
	visit, err := s.repo.GetLatestVisit(memberID)
	if err != nil {
		if err.Error() == domain.ErrNotInDB {
			return errors.New(domain.ErrIsNotInGym)
		}
		return err
	}
	if !visit.LeftAt.IsZero() {
		return errors.New(domain.ErrIsNotInGym)
	}

	return s.repo.EndVisit(visit.ID)
}
