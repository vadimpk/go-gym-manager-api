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

func (r *MemberRepo) Create(input domain.MemberCreateInput) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, phone_number, joined_at) VALUES ($1, $2, $3, $4) RETURNING id", membersTable)
	row := r.db.QueryRowx(query, input.FirstName, input.LastName, input.PhoneNumber, "NOW()")

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

func (r *MemberRepo) Update(id int, input domain.MemberUpdateInput) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, last_name = $2, phone_number = $3 WHERE id = $4", membersTable)
	_, err := r.db.Exec(query, input.FirstName, input.LastName, input.PhoneNumber, id)
	return err
}
func (r *MemberRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", membersTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MemberRepo) SetMembership(memberID int, membershipID int, expiresAt time.Time) error {
	query := fmt.Sprintf("INSERT INTO %s (membership_id, member_id, membership_expiration) VALUES ($1, $2, $3)", membersMembershipsTable)
	_, err := r.db.Exec(query, membershipID, memberID, expiresAt)
	return err
}

func (r *MemberRepo) UpdateMembership(memberID int, membershipID int, expiresAt time.Time) error {
	query := fmt.Sprintf("UPDATE %s SET membership_id = $1, membership_expiration = $2 WHERE member_id = $3", membersMembershipsTable)
	_, err := r.db.Exec(query, membershipID, expiresAt, memberID)
	return err
}

func (r *MemberRepo) GetMembership(memberID int) (domain.MembersMembership, error) {
	var membership domain.MembersMembership
	query := fmt.Sprintf("SELECT * FROM %s WHERE member_id = $1", membersMembershipsTable)
	err := r.db.Get(&membership, query, memberID)
	return membership, err
}

func (r *MemberRepo) DeleteMembership(memberID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = $1", membersMembershipsTable)
	_, err := r.db.Exec(query, memberID)
	return err
}

func (r *MemberRepo) GetLatestVisit(memberID int) (domain.MembersVisits, error) {
	var visit domain.MembersVisits
	query := fmt.Sprintf("SELECT * FROM %s WHERE member_id = $1 ORDER BY arrived_at DESC LIMIT 1", membersVisitsTable)
	err := r.db.Get(&visit, query, memberID)
	return visit, err
}

func (r *MemberRepo) SetNewVisit(memberID int, managerID int) error {
	// you should optimize data so that it doesn't have null values when working with go
	nullTime := new(domain.TrainersVisits).LeftAt
	query := fmt.Sprintf("INSERT INTO %s (manager_id, member_id, arrived_at, left_at) VALUES ($1, $2, $3, $4)", membersVisitsTable)
	_, err := r.db.Exec(query, managerID, memberID, "NOW()", nullTime)
	return err
}

func (r *MemberRepo) EndVisit(id int) error {
	query := fmt.Sprintf("UPDATE %s SET left_at = $1 WHERE id = $2", membersVisitsTable)
	_, err := r.db.Exec(query, "NOW()", id)
	return err
}
