package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rbueno/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuit struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (s *ClientDBTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	s.ClientDB = NewClientDB(db)
}

func (s *ClientDBTestSuit) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuit))
}

func (s *ClientDBTestSuit) TestSave() {
	client, _ := entity.NewClient("richard", "rich@email.com")
	err := s.ClientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuit) TestGet() {
	client, _ := entity.NewClient("richard", "rich@email.com")
	s.ClientDB.Save(client)

	clientDB, err := s.ClientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
