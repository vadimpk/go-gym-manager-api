package service

import (
	"errors"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
)

type TrainersService struct {
	repo repository.Trainers
}

func NewTrainersService(repo repository.Trainers) *TrainersService {
	return &TrainersService{repo: repo}
}

func (s *TrainersService) CreateNew(input domain.TrainerCreateInput) (int, error) {
	return s.repo.Create(input)
}

func (s *TrainersService) GetByID(id int) (domain.Trainer, error) {
	return s.repo.GetByID(id)
}

func (s *TrainersService) UpdateByID(id int, input domain.TrainerUpdateInput) error {
	trainer, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if input.FirstName == "" {
		input.FirstName = trainer.FirstName
	}
	if input.LastName == "" {
		input.LastName = trainer.LastName
	}
	if input.PhoneNumber == "" {
		input.PhoneNumber = trainer.PhoneNumber
	}
	if input.Email == "" {
		input.Email = trainer.Email
	}
	if input.Description == "" {
		input.Description = trainer.Description
	}
	if input.Price == 0 {
		input.Price = trainer.Price
	}

	return s.repo.Update(id, input)
}

func (s *TrainersService) DeleteByID(id int) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TrainersService) SetNewVisit(trainerID int, managerID int) error {
	_, err := s.repo.GetByID(trainerID)
	if err != nil {
		return err
	}

	visit, err := s.repo.GetLatestVisit(trainerID)
	if err != nil {
		if err.Error() == domain.ErrNotInDB {
			return s.repo.SetNewVisit(trainerID, managerID)
		}
		return err
	}
	if visit.LeftAt.IsZero() {
		return errors.New(domain.ErrStillInGym)
	}
	return s.repo.SetNewVisit(trainerID, managerID)
}

func (s *TrainersService) EndVisit(trainerID int) error {
	_, err := s.repo.GetByID(trainerID)
	if err != nil {
		return err
	}

	visit, err := s.repo.GetLatestVisit(trainerID)
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
