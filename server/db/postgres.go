package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"

	"github.com/gordonrehling2/certavs/config"
)

type IDB interface {
	Connect() *sql.DB
	Query(string, ...interface{}) (*sql.Rows, error)
}

type postgresDB struct {
	DBConfig config.Config
	db *sql.DB
}

func NewPostgresDB (config config.Config) *postgresDB {
	return &postgresDB{
		config,
		nil,
	}
}

func (p *postgresDB) Connect() *sql.DB {
	var err error

	p.db, err = sql.Open("postgres", p.DBConfig.DB.BuildConnectionURL())
	if err != nil {
		log.Fatalf("Connect failed %+v", err)
	}
	return p.db
}

func (p *postgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	var err error
	var rows *sql.Rows

	if args == nil {
		rows, err = p.db.Query(query)
	} else {
		rows, err = p.db.Query(query, args)
	}

	if err != nil {
		log.Fatalf("Query failed %+v", err)
	}
	return rows, err
}


