package postgres_gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func GetConnectionUri() string {
	return os.Getenv("POSTGRES_CONNECTION_URI")
}

type Postgres struct {
	Db *gorm.DB
}

func New(connectionUri string) (*Postgres, error) {
	db, err := gorm.Open("postgres_gorm", connectionUri)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}

func (p *Postgres) Close() {
	if err := p.Db.Close(); err != nil {
		fmt.Printf("could not close postgres connection: %s", err)
	}
}
