package database

import (
	"database/sql"
	"github.com/gaspartv/go-fullcycle-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDb *TransactionDb
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	_, err = db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE clients (id VARCHAT(255), name VARCHAT(255), email VARCHAT(255), created_at DATE)")
	s.Nil(err)

	_, err = db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	s.Nil(err)

	client1, err := entity.NewClient("John Doe", "j@j.com")
	s.Nil(err)
	s.client1 = client1

	client2, err := entity.NewClient("Jane Doe", "k@k.com")
	s.Nil(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(client1)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDb = NewTransactionDb(db)
}

func (s *TransactionDbTestSuite) TearDownSuite() {
	defer func(db *sql.DB) {
		err := db.Close()
		s.Nil(err)
	}(s.db)

	_, err := s.db.Exec("DROP TABLE transactions")
	s.Nil(err)

	_, err = s.db.Exec("DROP TABLE accounts")
	s.Nil(err)

	_, err = s.db.Exec("DROP TABLE clients")
	s.Nil(err)
}

func TestTransactionDbTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDb.Create(transaction)
	s.Nil(err)
}
