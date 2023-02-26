package psql

import (
	"database/sql"
	"fmt"

	"github.com/deed-labs/gittips/bot/internal/repository"
	_ "github.com/lib/pq"
)

type Database struct {
	dsn string
	db  *sql.DB

	ownersStorage   *ownersStorage
	bountiesStorage *bountiesStorage
}

func New(dsn string) (*Database, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	d := &Database{
		dsn:             dsn,
		db:              conn,
		ownersStorage:   &ownersStorage{db: conn},
		bountiesStorage: &bountiesStorage{db: conn},
	}

	return d, nil
}

func (d *Database) Owners() repository.OwnersRepository {
	return d.ownersStorage
}

func (d *Database) Bounties() repository.BountiesRepository {
	return d.bountiesStorage
}
