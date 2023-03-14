package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
)

type ownersStorage struct {
	db *sql.DB
}

func (s *ownersStorage) Get(ctx context.Context, ownerID int64) (*entity.Owner, error) {
	query := `SELECT gh_id, login, name, url, avatar_url, type, twitter_username, wallet_address FROM owners WHERE gh_id=$1`

	rows, err := s.db.QueryContext(ctx, query, ownerID)
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
		&owner.Name,
		&owner.URL,
		&owner.AvatarURL,
		&owner.Type,
		&owner.TwitterUsername,
		&owner.WalletAddress,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return owner, nil
}

func (s *ownersStorage) GetByWalletAddress(ctx context.Context, address string) (*entity.Owner, error) {
	query := `SELECT gh_id, login, name, url, avatar_url, type, twitter_username, wallet_address FROM owners WHERE wallet_address=$1`

	rows, err := s.db.QueryContext(ctx, query, address)
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
		&owner.Name,
		&owner.URL,
		&owner.AvatarURL,
		&owner.Type,
		&owner.TwitterUsername,
		&owner.WalletAddress,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return owner, nil
}

func (s *ownersStorage) Save(ctx context.Context, owner *entity.Owner) error {
	query := `INSERT INTO owners (
                      gh_id, login, name, url, avatar_url, type, twitter_username
                ) VALUES (
                          $1, $2, $3, $4, $5, $6, $7
                )
		ON CONFLICT (gh_id) DO UPDATE 
		SET login = excluded.login,
		    url = excluded.url,
		    avatar_url = excluded.avatar_url,
		    type = excluded.type,
		    twitter_username = excluded.twitter_username;
	`

	_, err := s.db.ExecContext(ctx, query, owner.ID, owner.Login, owner.Name, owner.URL, owner.AvatarURL,
		owner.Type, owner.TwitterUsername)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (s *ownersStorage) SetWalletAddress(ctx context.Context, ownerId int64, walletAddress string) error {
	query := `UPDATE owners SET wallet_address = $1 WHERE gh_id = $2`

	_, err := s.db.ExecContext(ctx, query, walletAddress, ownerId)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
