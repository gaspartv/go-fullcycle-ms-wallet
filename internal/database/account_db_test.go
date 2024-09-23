package database

import (
	"database/sql"
	"github.com/gaspartv/go-fullcycle-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDbTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	_, err = db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE clients (id VARCHAT(255), name VARCHAT(255), email VARCHAT(255), created_at DATE)")
	s.Nil(err)

	s.accountDB = NewAccountDb(db)
	s.client, _ = entity.NewClient("John Doe", "j@j.com")
}

func (s *AccountDbTestSuite) TearDownSuite() {
	err := s.db.Close()
	s.Nil(err)

	_, err = s.db.Exec("DROP TABLE accounts")
	s.Nil(err)

	_, err = s.db.Exec("DROP TABLE clients")
	s.Nil(err)
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestFindById() {
	s.db.Exec(
		"INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)

	accountDb, err := s.accountDB.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDb.ID)
	s.Equal(account.Client.ID, accountDb.Client.ID)
	s.Equal(account.Balance, accountDb.Balance)
	s.Equal(account.Client.Name, accountDb.Client.Name)
	s.Equal(account.Client.Email, accountDb.Client.Email)
}
