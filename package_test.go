package dao_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	// load PostgreSQL driver
	_ "github.com/lib/pq"

	"github.com/steenzout/go-dao"
	"github.com/steenzout/go-dao/mock"
	"github.com/stretchr/testify/suite"
)


var db *sql.DB

func init() {
	var err error

	host := "127.0.0.1"
	database := "travis_ci_test"
	port := 5432
	user := "postgres"

	dbinfo := fmt.Sprintf(
		"user=%s host=%s port=%d dbname=%s sslmode=disable",
		user, host, port, database)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connection to %v was established", dbinfo)
}

type DatabaseTestSuite struct {
	suite.Suite
	manager dao.Manager
	factory dao.Factory
}

func (s *DatabaseTestSuite) SetupSuite() {

	ds := &dao.DataSource{
		DB: db,
		Name: "travis_ci_test",
	}

	s.manager = dao.NewBaseManager()
	s.manager.RegisterDataSource(ds)

	s.factory = mock.NewFactory(ds)
	s.manager.RegisterFactory(s.factory)
}

// TestPackage test suite for the dao package.
func TestPackage(t *testing.T) {
	suite.Run(t, new(DAOTestSuite))
	suite.Run(t, new(HandlerTestSuite))
	suite.Run(t, new(ManagerTestSuite))
}
