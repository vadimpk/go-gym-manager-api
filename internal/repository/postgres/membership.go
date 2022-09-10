package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
)

type MembershipRepo struct {
	db *sqlx.DB
}

func NewMembershipRepo(db *sqlx.DB) *MembershipRepo {
	return &MembershipRepo{db: db}
}

func (r *MembershipRepo) Create(input domain.MembershipCreateInput) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (short_name, description, price, duration) VALUES ($1, $2, $3, $4) RETURNING id", membershipsTable)
	row := r.db.QueryRowx(query, input.ShortName, input.Description, input.Price, input.Duration)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MembershipRepo) GetByID(id int) (domain.Membership, error) {
	var membership domain.Membership
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", membershipsTable)
	err := r.db.Get(&membership, query, id)
	return membership, err
}

func (r *MembershipRepo) Update(id int, input domain.MembershipUpdateInput) error {
	query := fmt.Sprintf("UPDATE %s SET short_name = $1, description = $2, price = $3, duration = $4 WHERE id = $5", membershipsTable)
	_, err := r.db.Exec(query, input.ShortName, input.Description, input.Price, input.Duration, id)
	return err
}

func (r *MembershipRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", membershipsTable)
	_, err := r.db.Exec(query, id)
	return err
}
