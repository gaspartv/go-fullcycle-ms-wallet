package database

import (
	"database/sql"
	"github.com/gaspartv/go-fullcycle-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type ClientDbTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *ClientDb
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	_, err = db.Exec("CREATE TABLE clients (id VARCHAT(255), name VARCHAT(255), email VARCHAT(255), created_at DATE)")
	s.Nil(err)

	s.clientDb = NewClientDb(db)
}

func (s *ClientDbTestSuite) TearDownSuite() {
	err := s.db.Close()
	s.Nil(err)

	_, err = s.db.Exec("DROP TABLE clients")
	s.Nil(err)
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestSave() {
	client := &entity.Client{ID: "1", Name: "John Doe", Email: "j@j.com"}
	err := s.clientDb.Save(client)
	s.Nil(err)
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	err := s.clientDb.Save(client)
	s.Nil(err)

	clientDb, err := s.clientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}
