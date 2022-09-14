package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	v1 "github.com/vadimpk/go-gym-manager-api/internal/delivery/http/v1"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/internal/repository/postgres"
	"github.com/vadimpk/go-gym-manager-api/internal/service"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
	"testing"
	"time"
)

type APITestSuite struct {
	suite.Suite

	db       *sqlx.DB
	handlers *v1.Handler
	services *service.Services
	repos    *repository.Repositories
	router   *gin.Engine

	tokenManager auth.TokenManager
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	cfg := config.Config{
		DB: config.PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "lz921skm0001p",
			DBName:   "testing",
			SSLMode:  "disable",
		},
		Auth: config.AuthConfig{
			AccessTokenTTL:  time.Hour * 2,
			RefreshTokenTTL: time.Hour * 720,
			SigningKey:      "testkey",
			PasswordSalt:    "",
		},
	}

	if db, err := postgres.NewPostgresDB(&cfg); err != nil {
		s.FailNow("error when connecting to db", err)
	} else {
		s.db = db
	}

	s.tokenManager, _ = auth.NewManager(cfg.Auth.SigningKey)

	s.repos = repository.NewRepositories(s.db)
	s.services = service.NewServices(&cfg, s.tokenManager, s.repos)
	s.handlers = v1.NewHandler(s.services, s.tokenManager)

	s.router = gin.New()
	s.handlers.Init(s.router.Group("/api"))

	s.populateDB()
}

func (s *APITestSuite) populateDB() {
	_, err := s.db.Exec(createManagerQuery)
	if err != nil {
		s.FailNow("cannot create manager", err)
	}
}

func (s *APITestSuite) clearTables(tables []string) {
	for _, t := range tables {
		// delete data
		_, err := s.db.Exec(fmt.Sprintf(deleteDataQuery, t))
		if err != nil {
			s.FailNow("cannot delete data from table", err)
		}

		// reset sequence
		_, err = s.db.Exec(fmt.Sprintf(resetSequenceQuery, t, t))
		if err != nil {
			s.FailNow("cannot delete data from table", err)
		}
	}
}
