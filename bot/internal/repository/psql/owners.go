package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/deed-labs/openroll/bot/internal/entity"
	"github.com/deed-labs/openroll/bot/internal/repository"
)

type ownersStorage struct {
	db *sql.DB
}

func (s *ownersStorage) Get(ctx context.Context, ownerID int64) (*entity.Owner, error) {
	query := `SELECT gh_id, login, url, avatar_url, type, twitter_username FROM owners WHERE gh_id=$1`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	if !rows.Next() {
		return nil, repository.ErrNotFound
	}

	owner := new(entity.Owner)
	if err := rows.Scan(
		&owner.ID,
		&owner.Login,
		&owner.URL,
		&owner.AvatarURL,
		&owner.Type,
		&owner.TwitterUsername,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return owner, nil
}

func (s *ownersStorage) Save(ctx context.Context, owner *entity.Owner) error {
	query := `INSERT INTO owners (
                      gh_id, login, url, avatar_url, type, twitter_username
                ) VALUES (
                          $1, $2, $3, $4, $5, $6
                )
	`

	_, err := s.db.ExecContext(ctx, query, owner.ID, owner.Login, owner.URL, owner.AvatarURL,
		owner.Type, owner.TwitterUsername)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
