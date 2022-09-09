package postgres

import (
	"context"
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

func (r *ManagerRepo) GetByCredentials(ctx context.Context, input domain.SignInInput) (domain.Manager, error) {
	var manager domain.Manager
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_number=$1 and password=$2", managersTable)
	err := r.db.Get(&manager, query, input.PhoneNumber, input.Password)
	return manager, err
}

func (r *ManagerRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Manager, error) {
	return domain.Manager{}, nil
}

func (r *ManagerRepo) SetSession(ctx context.Context, managerID int, session domain.Session) error {
	return nil
}
