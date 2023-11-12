package db

import (
	"database/sql"
	"github.com/Wintersunner/xplor/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load test config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource())
	if err != nil {
		log.Fatal("Cannot connect to test database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
