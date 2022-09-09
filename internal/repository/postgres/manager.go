package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
)

type ManagerRepo struct {
	db *sqlx.DB
}

func NewManagerRepo(db *sqlx.DB) *ManagerRepo {
	return &ManagerRepo{db: db}
}

func (r *ManagerRepo) GetByCredentials(input domain.SignInInput) (domain.Manager, error) {
	var manager domain.Manager
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_number=$1 and password=$2", managersTable)
	err := r.db.Get(&manager, query, input.PhoneNumber, input.Password)
	return manager, err
}

func (r *ManagerRepo) GetByRefreshToken(refreshToken string) (int, error) {
	var session domain.Session
	query := fmt.Sprintf("SELECT * FROM %s WHERE refresh_token=$1", sessionsTable)
	err := r.db.Get(&session, query, refreshToken)
	return session.ManagerID, err
}

func (r *ManagerRepo) SetSession(managerID int, session domain.Session) error {

	// get current session
	var s domain.Session
	query := fmt.Sprintf("SELECT id FROM %s WHERE manager_id=$1", sessionsTable)
	err := r.db.Get(&s, query, managerID)

	// if session doesn't exist, create new one
	if s.ID == 0 {
		query := fmt.Sprintf("INSERT INTO %s (refresh_token, expires_at, manager_id) VALUES ($1, $2, $3)", sessionsTable)
		_, err := r.db.Exec(query, session.RefreshToken, session.ExpiresAt, managerID)
		return err
	}

	// if exists, update it
	query = fmt.Sprintf("UPDATE %s SET refresh_token = $1, expires_at = $2 WHERE manager_id = $3", sessionsTable)
	_, err = r.db.Exec(query, session.RefreshToken, session.ExpiresAt, managerID)
	return err
}
