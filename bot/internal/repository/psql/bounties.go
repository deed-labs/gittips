package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"

	"github.com/deed-labs/gittips/bot/internal/entity"
)

type bountiesStorage struct {
	db *sql.DB
}

func (s *bountiesStorage) GetAll(ctx context.Context) ([]*entity.BountyWithOwner, error) {
	query := `SELECT 
    	    	bounties.owner_gh_id, bounties.title, bounties.url, bounties.reward,
    	    	owners.login, owners.url, owners.avatar_url, owners.type FROM bounties, owners 
    	WHERE bounties.owner_gh_id = owners.gh_id AND bounties.closed = FALSE `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*entity.BountyWithOwner{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	bounties := make([]*entity.BountyWithOwner, 0)

	for rows.Next() {
		var reward int64

		bounty := new(entity.BountyWithOwner)
		if err := rows.Scan(
			&bounty.OwnerID,
			&bounty.Title,
			&bounty.URL,
			&reward,
			&bounty.OwnerLogin,
			&bounty.OwnerURL,
			&bounty.OwnerAvatarURL,
			&bounty.OwnerType,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		bounty.Reward = big.NewInt(reward)
		bounties = append(bounties, bounty)
	}

	return bounties, nil
}

func (s *bountiesStorage) GetByOwnerId(ctx context.Context, ownerId int64) ([]*entity.Bounty, error) {
	query := `SELECT  gh_id, owner_gh_id, title, url, reward, closed
    	    FROM bounties WHERE owner_gh_id = $1`

	rows, err := s.db.QueryContext(ctx, query, ownerId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []*entity.Bounty{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	bounties := make([]*entity.Bounty, 0)

	for rows.Next() {
		var reward int64

		bounty := new(entity.Bounty)
		if err := rows.Scan(
			&bounty.ID,
			&bounty.OwnerID,
			&bounty.Title,
			&bounty.URL,
			&reward,
			&bounty.Closed,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		bounty.Reward = big.NewInt(reward)
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

	_, err := s.db.ExecContext(ctx, query, bounty.ID, bounty.OwnerID, bounty.Title, bounty.URL, bounty.Reward.String())
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

func (s *bountiesStorage) SetReward(ctx context.Context, bountyId int64, value *big.Int) error {
	query := `UPDATE bounties SET reward = $1 WHERE gh_id = $2`

	_, err := s.db.ExecContext(ctx, query, value.String(), bountyId)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (s *bountiesStorage) SetClosed(ctx context.Context, bountyId int64, closed bool) error {
	query := `UPDATE bounties SET closed = $1 WHERE gh_id = $2`

	_, err := s.db.ExecContext(ctx, query, closed, bountyId)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
