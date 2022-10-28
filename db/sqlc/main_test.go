package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jmbustos1/simplebank/util"
	_ "github.com/lib/pq"
)

/*
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	//dbSource = "postgresql://root:secret@localhost:5433/simple_bank2?sslmode=disable"
)
*/

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	var err error
	var config util.Config

	config, err = util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
