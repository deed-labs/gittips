package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/deed-labs/gittips/bot/internal/entity"
)

type bountiesStorage struct {
	db *sql.DB
}

func (s *bountiesStorage) GetAll(ctx context.Context) ([]*entity.Bounty, error) {
	query := `SELECT 
    	    	bounties.owner_gh_id, bounties.title, bounties.url, bounties.reward,
    	    	owners.login, owners.url, owners.avatar_url, owners.type FROM bounties, owners 
    	WHERE bounties.owner_gh_id = owners.gh_id`

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
			&bounty.OwnerLogin,
			&bounty.OwnerURL,
			&bounty.OwnerAvatarURL,
			&bounty.OwnerType,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		bounties = append(bounties, bounty)
	}

	return bounties, nil
}

func (s *bountiesStorage) Save(ctx context.Context, bounty *entity.Bounty) error {
	query := `INSERT INTO bounties (
                    gh_id, owner_gh_id, title, url, reward
                ) VALUES (
                    $1, $2, $3, $4, $5
                )
	`

	_, err := s.db.ExecContext(ctx, query, bounty.ID, bounty.OwnerID, bounty.Title, bounty.URL, bounty.Reward)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (s *bountiesStorage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM bounties WHERE gh_id=$1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
