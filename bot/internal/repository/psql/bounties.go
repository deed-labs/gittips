package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/deed-labs/openroll/bot/internal/entity"
)

type bountiesStorage struct {
	db *sql.DB
}

func (s *bountiesStorage) GetAll(ctx context.Context) ([]*entity.Bounty, error) {
	query := `SELECT owner_id, title, url, reward FROM bounties`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*entity.Bounty{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	bounties := make([]*entity.Bounty, 0)

	for rows.Next() {
		bounty := new(entity.Bounty)
		if err := rows.Scan(
			&bounty.OwnerID,
			&bounty.Title,
			&bounty.URL,
			&bounty.Reward,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		bounties = append(bounties, bounty)
	}

	return bounties, nil
}

func (s *bountiesStorage) Save(ctx context.Context, bounty *entity.Bounty) error {
	query := `INSERT INTO bounties (
                    owner_id, title, url, reward
                ) VALUES (
                    $1, $2, $3, $4
                )
	`

	_, err := s.db.ExecContext(ctx, query, bounty.OwnerID, bounty.Title, bounty.URL, bounty.Reward)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
