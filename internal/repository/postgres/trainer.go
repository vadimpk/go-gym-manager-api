package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
)

type TrainerRepo struct {
	db *sqlx.DB
}

func NewTrainerRepo(db *sqlx.DB) *TrainerRepo {
	return &TrainerRepo{db: db}
}

func (r *TrainerRepo) Create(input domain.TrainerCreateInput) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, phone_number, email, description, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", trainersTable)
	row := r.db.QueryRowx(query, input.FirstName, input.LastName, input.PhoneNumber, input.Email, input.Description, input.Price)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TrainerRepo) GetByID(id int) (domain.Trainer, error) {
	var trainer domain.Trainer
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", trainersTable)
	err := r.db.Get(&trainer, query, id)
	return trainer, err
}

func (r *TrainerRepo) Update(id int, input domain.TrainerUpdateInput) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, last_name = $2, phone_number = $3, email = $4, description = $5, price = $6 WHERE id = $7", trainersTable)
	_, err := r.db.Exec(query, input.FirstName, input.LastName, input.PhoneNumber, input.Email, input.Description, input.Price, id)
	return err
}

func (r *TrainerRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", membershipsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TrainerRepo) GetLatestVisit(trainerID int) (domain.TrainersVisits, error) {
	var visit domain.TrainersVisits
	query := fmt.Sprintf("SELECT * FROM %s WHERE trainer_id = $1 ORDER BY arrived_at DESC LIMIT 1", trainersVisitsTable)
	err := r.db.Get(&visit, query, trainerID)
	return visit, err
}

func (r *TrainerRepo) SetNewVisit(trainerID int, managerID int) error {
	nullTime := new(domain.TrainersVisits).LeftAt
	query := fmt.Sprintf("INSERT INTO %s (manager_id, trainer_id, arrived_at, left_at) VALUES ($1, $2, $3, $4)", trainersVisitsTable)
	_, err := r.db.Exec(query, managerID, trainerID, "NOW()", nullTime)
	return err
}

func (r *TrainerRepo) EndVisit(id int) error {
	query := fmt.Sprintf("UPDATE %s SET left_at = $1 WHERE id = $2", trainersVisitsTable)
	_, err := r.db.Exec(query, "NOW()", id)
	return err
}
