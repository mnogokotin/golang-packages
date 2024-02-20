package postgres

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func New(connectionUri string) (*Postgres, error) {
	db, err := sql.Open("postgres", connectionUri)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) Close() {
	if err := p.db.Close(); err != nil {
		fmt.Printf("could not close postgres connection: %s", err)
	}
}
