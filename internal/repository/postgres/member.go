package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"time"
)

type MemberRepo struct {
	db *sqlx.DB
}

func NewMemberRepo(db *sqlx.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

func (r *MemberRepo) Create(input domain.MemberCreate) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, phone_number, membership_id, joined_at, membership_expiration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", membersTable)
	row := r.db.QueryRowx(query, input.FirstName, input.LastName, input.PhoneNumber, input.MembershipID, "NOW()", "NOW()")

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MemberRepo) GetByID(id int) (domain.Member, error) {
	var member domain.Member
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", membersTable)
	err := r.db.Get(&member, query, id)
	return member, err
}

func (r *MemberRepo) GetByPhoneNumber(num string) (domain.Member, error) {
	var member domain.Member
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone_number = $1", membersTable)
	err := r.db.Get(&member, query, num)
	return member, err
}

func (r *MemberRepo) Update(id int, input domain.MemberUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, last_name = $2, phone_number = $3 WHERE id = $4", membersTable)
	_, err := r.db.Exec(query, input.FirstName, input.LastName, input.PhoneNumber, id)
	return err
}
func (r *MemberRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", membersTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MemberRepo) SetMembership(id int, membershipID int) error {
	query := fmt.Sprintf("UPDATE %s SET memberhip_id = $1 WHERE id = $2", membersTable)
	_, err := r.db.Exec(query, membershipID, id)
	return err
}
func (r *MemberRepo) DeleteMembership(id int) error {
	query := fmt.Sprintf("UPDATE %s SET memberhip_id = $1, expires_at = $2 WHERE id = $2", membersTable)
	_, err := r.db.Exec(query, 0, time.Now(), id)
	return err
}
