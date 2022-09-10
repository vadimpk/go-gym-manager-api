package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vadimpk/go-gym-manager-api/internal/config"
)

const (
	managersTable           = "managers"
	sessionsTable           = "sessions"
	membersTable            = "members"
	membershipsTable        = "memberships"
	trainersTable           = "trainers"
	membersMembershipsTable = "members_memberships"
	trainersVisitsTable     = "trainers_visits"
	membersVisitsTable      = "members_visits"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
